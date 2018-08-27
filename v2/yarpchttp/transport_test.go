// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package yarpchttp

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/yarpc/internal/testtime"
	"go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpctest"
)

// NoJitter is a transport option only available in tests, to disable jitter
// between connection attempts.
func NoJitter() TransportOption {
	return func(options *transportOptions) {
		options.jitter = func(n int64) int64 {
			return n
		}
	}
}

type peerExpectation struct {
	id          string
	subscribers []string
}

func createPeerIdentifierMap(ids []string) map[string]yarpc.Identifier {
	pids := make(map[string]yarpc.Identifier, len(ids))
	for _, id := range ids {
		pids[id] = &testIdentifier{id}
	}
	return pids
}

func TestTransport(t *testing.T) {
	type testStruct struct {
		msg string

		// identifiers defines all the Identifiers that will be used in
		// the actions up from so they can be generated and passed as deps
		identifiers []string

		// subscriberDefs defines all the Subscribers that will be used in
		// the actions up from so they can be generated and passed as deps
		subscriberDefs []yarpctest.SubscriberDefinition

		// actions are the actions that will be applied against the transport
		actions []yarpctest.DialerAction

		// expectedPeers are a list of peers (and those peer's subscribers)
		// that are expected on the transport after the actions
		expectedPeers []peerExpectation
	}
	tests := []testStruct{
		{
			msg:         "one retain",
			identifiers: []string{"i1"},
			subscriberDefs: []yarpctest.SubscriberDefinition{
				{ID: "s1"},
			},
			actions: []yarpctest.DialerAction{
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s1", ExpectedPeerID: "i1"},
			},
			expectedPeers: []peerExpectation{
				{id: "i1", subscribers: []string{"s1"}},
			},
		},
		{
			msg:         "one retain one release",
			identifiers: []string{"i1"},
			subscriberDefs: []yarpctest.SubscriberDefinition{
				{ID: "s1"},
			},
			actions: []yarpctest.DialerAction{
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s1", ExpectedPeerID: "i1"},
				yarpctest.ReleaseAction{InputIdentifierID: "i1", InputSubscriberID: "s1"},
			},
		},
		{
			msg:         "three retains",
			identifiers: []string{"i1"},
			subscriberDefs: []yarpctest.SubscriberDefinition{
				{ID: "s1"},
				{ID: "s2"},
				{ID: "s3"},
			},
			actions: []yarpctest.DialerAction{
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s1", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s2", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s3", ExpectedPeerID: "i1"},
			},
			expectedPeers: []peerExpectation{
				{id: "i1", subscribers: []string{"s1", "s2", "s3"}},
			},
		},
		{
			msg:         "three retains one release",
			identifiers: []string{"i1"},
			subscriberDefs: []yarpctest.SubscriberDefinition{
				{ID: "s1"},
				{ID: "s2r"},
				{ID: "s3"},
			},
			actions: []yarpctest.DialerAction{
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s1", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s2r", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s3", ExpectedPeerID: "i1"},
				yarpctest.ReleaseAction{InputIdentifierID: "i1", InputSubscriberID: "s2r"},
			},
			expectedPeers: []peerExpectation{
				{id: "i1", subscribers: []string{"s1", "s3"}},
			},
		},
		{
			msg:         "three retains, three release",
			identifiers: []string{"i1"},
			subscriberDefs: []yarpctest.SubscriberDefinition{
				{ID: "s1"},
				{ID: "s2"},
				{ID: "s3"},
			},
			actions: []yarpctest.DialerAction{
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s1", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s2", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s3", ExpectedPeerID: "i1"},
				yarpctest.ReleaseAction{InputIdentifierID: "i1", InputSubscriberID: "s1"},
				yarpctest.ReleaseAction{InputIdentifierID: "i1", InputSubscriberID: "s2"},
				yarpctest.ReleaseAction{InputIdentifierID: "i1", InputSubscriberID: "s3"},
			},
		},
		{
			msg:         "no retains one release",
			identifiers: []string{"i1"},
			subscriberDefs: []yarpctest.SubscriberDefinition{
				{ID: "s1"},
			},
			actions: []yarpctest.DialerAction{
				yarpctest.ReleaseAction{
					InputIdentifierID: "i1",
					InputSubscriberID: "s1",
					ExpectedErrType:   yarpc.ErrTransportHasNoReferenceToPeer{},
				},
			},
		},
		{
			msg:         "one retains, one release (from different subscriber)",
			identifiers: []string{"i1"},
			subscriberDefs: []yarpctest.SubscriberDefinition{
				{ID: "s1"},
				{ID: "s2"},
			},
			actions: []yarpctest.DialerAction{
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s1", ExpectedPeerID: "i1"},
				yarpctest.ReleaseAction{
					InputIdentifierID: "i1",
					InputSubscriberID: "s2",
					ExpectedErrType:   yarpc.ErrPeerHasNoReferenceToSubscriber{},
				},
			},
			expectedPeers: []peerExpectation{
				{id: "i1", subscribers: []string{"s1"}},
			},
		},
		{
			msg:         "multi peer retain/release",
			identifiers: []string{"i1", "i2", "i3", "i4r", "i5r"},
			subscriberDefs: []yarpctest.SubscriberDefinition{
				{ID: "s1"},
				{ID: "s2"},
				{ID: "s3"},
				{ID: "s4"},
				{ID: "s5rnd"},
				{ID: "s6rnd"},
				{ID: "s7rnd"},
			},
			actions: []yarpctest.DialerAction{
				// Retains/Releases of i1 (Retain/Release the random peers at the end)
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s5rnd", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s6rnd", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s1", ExpectedPeerID: "i1"},
				yarpctest.RetainAction{InputIdentifierID: "i1", InputSubscriberID: "s2", ExpectedPeerID: "i1"},
				yarpctest.ReleaseAction{InputIdentifierID: "i1", InputSubscriberID: "s5rnd"},
				yarpctest.ReleaseAction{InputIdentifierID: "i1", InputSubscriberID: "s6rnd"},

				// Retains/Releases of i2 (Retain then Release then Retain again)
				yarpctest.RetainAction{InputIdentifierID: "i2", InputSubscriberID: "s2", ExpectedPeerID: "i2"},
				yarpctest.RetainAction{InputIdentifierID: "i2", InputSubscriberID: "s3", ExpectedPeerID: "i2"},
				yarpctest.ReleaseAction{InputIdentifierID: "i2", InputSubscriberID: "s2"},
				yarpctest.ReleaseAction{InputIdentifierID: "i2", InputSubscriberID: "s3"},
				yarpctest.RetainAction{InputIdentifierID: "i2", InputSubscriberID: "s2", ExpectedPeerID: "i2"},
				yarpctest.RetainAction{InputIdentifierID: "i2", InputSubscriberID: "s3", ExpectedPeerID: "i2"},

				// Retains/Releases of i3 (Retain/Release unrelated sub, then retain two)
				yarpctest.RetainAction{InputIdentifierID: "i3", InputSubscriberID: "s7rnd", ExpectedPeerID: "i3"},
				yarpctest.ReleaseAction{InputIdentifierID: "i3", InputSubscriberID: "s7rnd"},
				yarpctest.RetainAction{InputIdentifierID: "i3", InputSubscriberID: "s3", ExpectedPeerID: "i3"},
				yarpctest.RetainAction{InputIdentifierID: "i3", InputSubscriberID: "s4", ExpectedPeerID: "i3"},

				// Retain/Release i4r on random subscriber
				yarpctest.RetainAction{InputIdentifierID: "i4r", InputSubscriberID: "s5rnd", ExpectedPeerID: "i4r"},
				yarpctest.ReleaseAction{InputIdentifierID: "i4r", InputSubscriberID: "s5rnd"},

				// Retain/Release i5r on already used subscriber
				yarpctest.RetainAction{InputIdentifierID: "i5r", InputSubscriberID: "s3", ExpectedPeerID: "i5r"},
				yarpctest.ReleaseAction{InputIdentifierID: "i5r", InputSubscriberID: "s3"},
			},
			expectedPeers: []peerExpectation{
				{id: "i1", subscribers: []string{"s1", "s2"}},
				{id: "i2", subscribers: []string{"s2", "s3"}},
				{id: "i3", subscribers: []string{"s3", "s4"}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.msg, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			trans := NewTransport()

			deps := yarpctest.DialerDeps{
				PeerIdentifiers: createPeerIdentifierMap(tt.identifiers),
				Subscribers:     yarpctest.CreateSubscriberMap(mockCtrl, tt.subscriberDefs),
			}
			yarpctest.ApplyDialerActions(t, trans, tt.actions, deps)

			assert.Len(t, trans.peers, len(tt.expectedPeers))
			for _, expectedPeerNode := range tt.expectedPeers {
				p, ok := trans.peers[expectedPeerNode.id]
				assert.True(t, ok)

				if assert.NotNil(t, p) {
					assert.Equal(t, expectedPeerNode.id, p.Identifier())

					// We can't look at the subscribers directly so we'll
					// attempt to remove subscribers and be sure that it
					// doesn't error
					assert.Len(t, expectedPeerNode.subscribers, p.NumSubscribers())
					for _, sub := range expectedPeerNode.subscribers {
						err := p.Unsubscribe(deps.Subscribers[sub])
						assert.NoError(t, err, "peer %s did not have reference to subscriber %s", p.Identifier(), sub)
					}
				}
			}
		})
	}
}

func TestTransportClient(t *testing.T) {
	trans := NewTransport()

	assert.NotNil(t, trans.client)
}

func TestTransportClientOpaqueOptions(t *testing.T) {
	// Unfortunately the KeepAlive is obfuscated in the client, so we can't really
	// assert this worked.
	trans := NewTransport(
		KeepAlive(testtime.Second),
		MaxIdleConns(100),
		MaxIdleConnsPerHost(10),
		IdleConnTimeout(1*time.Second),
		DisableCompression(),
		DisableKeepAlives(),
		ResponseHeaderTimeout(1*time.Second),
	)

	assert.NotNil(t, trans.client)
}

type testIdentifier struct {
	id string
}

func (i testIdentifier) Identifier() string {
	return i.id
}
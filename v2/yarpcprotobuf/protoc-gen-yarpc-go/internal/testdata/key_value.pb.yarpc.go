// Code generated by protoc-gen-yarpc-go
// source: key_value.proto
// DO NOT EDIT!

package keyvalue

import (
	context "context"

	proto "github.com/gogo/protobuf/proto"
	yarpc "go.uber.org/yarpc/v2/yarpc"
	yarpcprotobuf "go.uber.org/yarpc/v2/yarpcprotobuf"
)

// KeyValueClient is the KeyValue service's client interface.
type KeyValueClient interface {
	Foo(
		context.Context,
		*GetRequest,
		...yarpc.CallOption,
	) (*GetResponse, error)
	Bar(
		context.Context,
		...yarpc.CallOption,
	) (KeyValueBarClientStream, error)
	Baz(
		context.Context,
		*GetRequest,
		...yarpc.CallOption,
	) (KeyValueBazClientStream, error)
	Qux(
		context.Context,
		...yarpc.CallOption,
	) (KeyValueQuxClientStream, error)
}

// NewKeyValueClient builds a new YARPC client for the KeyValue service.
func NewKeyValueClient(c yarpc.Client, opts ...yarpcprotobuf.ClientOption) KeyValueClient {
	return &_KeyValueCaller{stream: yarpcprotobuf.NewStreamClient(c, "KeyValue", opts...)}
}

type _KeyValueCaller struct {
	stream yarpcprotobuf.StreamClient
}

func (c *_KeyValueCaller) Foo(ctx context.Context, req *GetRequest, opts ...yarpc.CallOption) (*GetResponse, error) {
	msg, err := c.stream.Call(ctx, "Foo", req, newKeyValueFooResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*GetResponse)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyKeyValueFooResponse, res)
	}
	return res, nil
}

func (c *_KeyValueCaller) Bar(ctx context.Context, opts ...yarpc.CallOption) (KeyValueBarClientStream, error) {
	s, err := c.stream.CallStream(ctx, "Bar", opts...)
	if err != nil {
		return err
	}
	return &_KeyValueBarClientStream{stream: s}, nil
}

func (c *_KeyValueCaller) Baz(ctx context.Context, req *GetRequest, opts ...yarpc.CallOption) (KeyValueBazClientStream, error) {
	s, err := c.stream.CallStream(ctx, "Baz", opts...)
	if err != nil {
		return err
	}
	if err := s.Send(req); err != nil {
		return nil, err
	}
	return &_KeyValueBazClientStream{stream: s}, nil
}

func (c *_KeyValueCaller) Qux(ctx context.Context, opts ...yarpc.CallOption) (KeyValueQuxClientStream, error) {
	s, err := c.stream.CallStream(ctx, "Qux", opts...)
	if err != nil {
		return nil, err
	}
	return &_KeyValueQuxClientStream{stream: s}, nil
}

// KeyValueBarClientStream is a streaming interface used in the KeyValue}Client interface.
type KeyValueBarClientStream interface {
	Context() context.Context
	Send(*GetRequest, ...yarpc.StreamOption) error
	CloseAndRecv(...yarpc.StreamOption) (*GetResponse, error)
}

// KeyValueBazClientStream is a streaming interface used in the KeyValue}Client interface.
type KeyValueBazClientStream interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*GetResponse, error)
	CloseSend(...yarpc.StreamOption) error
}

// KeyValueQuxClientStream is a streaming interface used in the KeyValue}Client interface.
type KeyValueQuxClientStream interface {
	Context() context.Context
	Send(*GetRequest, ...yarpc.StreamOption) error
	Recv(...yarpc.StreamOption) (*GetResponse, error)
	CloseSend(...yarpc.StreamOption) error
}

type _KeyValueBarClientStream struct {
	stream *yarpcprotobuf.ClientStream
}

func (c *_KeyValueBarClientStream) Context() context.Context {
	return c.stream.Context()
}

func (c *_KeyValueBarClientStream) Send(req *GetRequest, opts ...yarpc.StreamOption) error {
	return c.stream.Send(req, opts...)
}

func (c *_KeyValueBarClientStream) CloseAndRecv(opts ...yarpc.StreamOption) (*GetRequest, error) {
	if err := c.stream.Close(opts...); err != nil {
		return nil, err
	}
	msg, err := c.stream.Receive(newKeyValueBarResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*GetResponse)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyKeyValueBarResponse, msg)
	}
	return res, err
}

type _KeyValueBazClientStream struct {
	stream *yarpcprotobuf.ClientStream
}

func (c *_KeyValueBazClientStream) Context() context.Context {
	return c.stream.Context()
}

func (c *_KeyValueBazClientStream) Recv(opts ...yarpc.StreamOption) (*GetResponse, error) {
	msg, err := c.stream.Receive(newKeyValueBazResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*KeyValueBazResponse)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyKeyValueBazResponse, msg)
	}
	return res, nil
}

func (c *_KeyValueBazClientStream) CloseSend(opts ...yarpc.StreamOption) error {
	return c.stream.Close(opts...)
}

type _KeyValueQuxClientStream struct {
	stream *yarpcprotobuf.ClientStream
}

func (c *_KeyValueQuxClientStream) Context() context.Context {
	return c.stream.Context()
}

func (c *_KeyValueQuxClientStream) Send(req *GetRequest, opts ...yarpc.StreamOption) error {
	return c.stream.Send(req, opts...)
}

func (c *_KeyValueQuxClientStream) Recv(opts ...yarpc.StreamOption) (*GetResponse, error) {
	msg, err := c.stream.Receive(newKeyValueQuxResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*KeyValueQuxResponse)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyKeyValueQuxResponse, msg)
	}
	return res, nil
}

func (c *_KeyValueQuxClientStream) CloseSend(opts ...yarpc.StreamOption) error {
	return c.stream.Close(opts...)
}

// KeyValueServer is the KeyValue service's server interface.
type KeyValueServer interface {
	Foo(
		context.Context,
		*GetRequest,
	) (*GetResponse, error)
	Bar(
		KeyValueBarServerStream,
	) (KeyValueBarServerStream, error)
	Baz(
		*GetRequest,
		KeyValueBazServerStream,
	) error
	Qux(
		KeyValueQuxServerStream,
	) error
}

// BuildKeyValueProcedures constructs the YARPC procedures for the KeyValue service.
func BuildKeyValueProcedures(s KeyValueServer) []yarpc.Procedure {
	h := &_KeyValueHandler{server: s}
	return yarpcprotobuf.Procedures(
		yarpcprotobuf.ProceduresParams{
			Service: "KeyValue",
			Unary: []yarpcprotobuf.UnaryProceduresParams{
				{
					MethodName: "Foo",
					Handler: yarpcprotobuf.NewUnaryHandler{
						yarpcprotobuf.UnaryHandlerParams{
							Handle: h.Foo,
							Create: newKeyValueFooRequest(),
						},
					},
				},
			},
			Stream: []yarpcprotobuf.StreamProceduresParams{
				{
					MethodName: "Bar",
					Handler: yarpcprotobuf.NewStreamHandler{
						yarpcprotobuf.StreamHandlerParams{
							Handle: h.Bar,
						},
					},
				},
				{
					MethodName: "Baz",
					Handler: yarpcprotobuf.NewStreamHandler{
						yarpcprotobuf.StreamHandlerParams{
							Handle: h.Baz,
						},
					},
				},
				{
					MethodName: "Qux",
					Handler: yarpcprotobuf.NewStreamHandler{
						yarpcprotobuf.StreamHandlerParams{
							Handle: h.Qux,
						},
					},
				},
			},
		},
	)
}

type _KeyValueHandler struct {
	server KeyValueServer
}

func (h *_KeyValueHandler) Foo(ctx context.Context, m proto.Message) (proto.Message, error) {
	req, _ := m.(*GetRequest)
	if req == nil {
		return nil, yarpcprotobuf.CastError(_emptyKeyValueFooRequest, m)
	}
	return h.server.Foo(ctx, req)
}

func (h *_KeyValueHandler) Bar(s *yarpcprotobuf.ServerStream) error {
	res, err := h.server.Bar(&_KeyValueBarServerStream{server: s})
	if err != nil {
		return err
	}
	return s.Send(res)
}

func (h *_KeyValueHandler) Baz(s *yarpcprotobuf.ServerStream) error {
	recv, err := s.Receive(newKeyValueBazRequest)
	if err != nil {
		return err
	}
	req, _ := recv.(*KeyValueBazRequest)
	if req == nil {
		return yarpcprotobuf.CastError(_emptyKeyValueBazRequest, recv)
	}
	return h.server.Baz(req, &_KeyValueBazServerStream{server: s})
}

func (h *_KeyValueHandler) Qux(s *yarpcprotobuf.ServerStream) error {
	return h.server.Qux(&_KeyValueQuxServerStream{stream: s})
}

// KeyValueBarServerStream is a streaming interface used in the KeyValue}Server interface.
type KeyValueBarServerStream interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*GetRequest, error)
}

// KeyValueBazServerStream is a streaming interface used in the KeyValue}Server interface.
type KeyValueBazServerStream interface {
	Context() context.Context
	Send(*GetResponse, ...yarpc.StreamOption) error
}

// KeyValueQuxServerStream is a streaming interface used in the KeyValue}Server interface.
type KeyValueQuxServerStream interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*GetRequest, error)
	Send(*GetResponse, ...yarpc.StreamOption) error
}

type _KeyValueBarServerStream struct {
	stream *yarpcprotobuf.ServerStream
}

func (s *_KeyValueBarServerStream) Context() context.Context {
	return s.stream.Context()
}

func (s *_KeyValueBarServerStream) Recv(opts ...yarpc.StreamOption) (*GetRequest, error) {
	msg, err := s.stream.Receive(newKeyValueBarRequest, opts...)
	if err != nil {
		return nil, err
	}
	req, ok := msg.(*KeyValueBarRequest)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyKeyValueBarRequest, msg)
	}
	return req, nil
}

type _KeyValueBazServerStream struct {
	stream *yarpcprotobuf.ServerStream
}

func (s *_KeyValueBazServerStream) Context() context.Context {
	return s.stream.Context()
}

func (s *_KeyValueBazServerStream) Send(res *GetResponse, opts ...yarpc.StreamOption) error {
	return s.stream.Send(res, opts...)
}

type _KeyValueQuxServerStream struct {
	stream *yarpcprotobuf.ServerStream
}

func (s *_KeyValueQuxServerStream) Context() context.Context {
	return s.stream.Context()
}

func (s *_KeyValueQuxServerStream) Recv(opts ...yarpc.StreamOption) (*GetRequest, error) {
	msg, err := s.stream.Receive(newKeyValueQuxRequest, opts...)
	if err != nil {
		return nil, err
	}
	req, ok := msg.(*KeyValueQuxRequest)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyKeyValueQuxRequest, msg)
	}
	return req, nil
}

func (s *_KeyValueQuxServerStream) Send(res *GetResponse, opts ...yarpc.StreamOption) error {
	return s.stream.Send(res, opts...)
}

func newKeyValueFooRequest()  { return &GetRequest{} }
func newKeyValueFooResponse() { return &GetResponse{} }
func newKeyValueBarRequest()  { return &GetRequest{} }
func newKeyValueBarResponse() { return &GetResponse{} }
func newKeyValueBazRequest()  { return &GetRequest{} }
func newKeyValueBazResponse() { return &GetResponse{} }
func newKeyValueQuxRequest()  { return &GetRequest{} }
func newKeyValueQuxResponse() { return &GetResponse{} }

var (
	_emptyKeyValueFooRequest  = &GetRequest{}
	_emptyKeyValueFooResponse = &GetResponse{}
	_emptyKeyValueBarRequest  = &GetRequest{}
	_emptyKeyValueBarResponse = &GetResponse{}
	_emptyKeyValueBazRequest  = &GetRequest{}
	_emptyKeyValueBazResponse = &GetResponse{}
	_emptyKeyValueQuxRequest  = &GetRequest{}
	_emptyKeyValueQuxResponse = &GetResponse{}
)

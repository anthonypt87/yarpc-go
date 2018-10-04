// Code generated by protoc-gen-yarpc-go
// source: src/keyvalue/key_value.proto
// DO NOT EDIT!

package keyvaluepb

import (
	context "context"
	proto "github.com/gogo/protobuf/proto"
	fx "go.uber.org/fx"
	yarpc "go.uber.org/yarpc/v2"
	yarpcprotobuf "go.uber.org/yarpc/v2/yarpcprotobuf"
	common "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/testdata/gen/proto/src/common"
)

// StoreClient is the Store service's client interface.
type StoreClient interface {
	Get(
		context.Context,
		*common.GetRequest,
		...yarpc.CallOption,
	) (*common.GetResponse, error)
	Set(
		context.Context,
		*common.SetRequest,
		...yarpc.CallOption,
	) (*common.SetResponse, error)
}

// NewStoreClient builds a new YARPC client for the Store service.
func NewStoreClient(c yarpc.Client, opts ...yarpcprotobuf.ClientOption) StoreClient {
	return &_StoreClient{stream: yarpcprotobuf.NewStreamClient(c, "keyvalue.Store", opts...)}
}

type _StoreClient struct {
	stream yarpcprotobuf.StreamClient
}

var _ StoreClient = (*_StoreClient)(nil)

func (c *_StoreClient) Get(ctx context.Context, req *common.GetRequest, opts ...yarpc.CallOption) (*common.GetResponse, error) {
	msg, err := c.stream.Call(ctx, "Get", req, newStoreGetResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*common.GetResponse)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyStoreGetResponse, res)
	}
	return res, nil
}

func (c *_StoreClient) Set(ctx context.Context, req *common.SetRequest, opts ...yarpc.CallOption) (*common.SetResponse, error) {
	msg, err := c.stream.Call(ctx, "Set", req, newStoreSetResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*common.SetResponse)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyStoreSetResponse, res)
	}
	return res, nil
}

// StoreServer is the Store service's server interface.
type StoreServer interface {
	Get(
		context.Context,
		*common.GetRequest,
	) (*common.GetResponse, error)
	Set(
		context.Context,
		*common.SetRequest,
	) (*common.SetResponse, error)
}

// BuildStoreProcedures constructs the YARPC procedures for the Store service.
func BuildStoreProcedures(s StoreServer) []yarpc.Procedure {
	h := &_StoreServer{server: s}
	return yarpcprotobuf.Procedures(
		yarpcprotobuf.ProceduresParams{
			Service: "keyvalue.Store",
			Unary: []yarpcprotobuf.UnaryProceduresParams{
				{
					Method: "Get",
					Handler: yarpcprotobuf.NewUnaryHandler(
						yarpcprotobuf.UnaryHandlerParams{
							Handle: h.Get,
							Create: newStoreGetRequest,
						},
					),
				},
				{
					Method: "Set",
					Handler: yarpcprotobuf.NewUnaryHandler(
						yarpcprotobuf.UnaryHandlerParams{
							Handle: h.Set,
							Create: newStoreSetRequest,
						},
					),
				},
			},
			Stream: []yarpcprotobuf.StreamProceduresParams{},
		},
	)
}

type _StoreServer struct {
	server StoreServer
}

func (h *_StoreServer) Get(ctx context.Context, m proto.Message) (proto.Message, error) {
	req, _ := m.(*common.GetRequest)
	if req == nil {
		return nil, yarpcprotobuf.CastError(_emptyStoreGetRequest, m)
	}
	return h.server.Get(ctx, req)
}

func (h *_StoreServer) Set(ctx context.Context, m proto.Message) (proto.Message, error) {
	req, _ := m.(*common.SetRequest)
	if req == nil {
		return nil, yarpcprotobuf.CastError(_emptyStoreSetRequest, m)
	}
	return h.server.Set(ctx, req)
}

// FxStoreClientParams defines the parameters
// required to provide a StoreClient into an
// Fx application.
type FxStoreClientParams struct {
	fx.In

	Client yarpc.Client
}

// FxStoreClientResult provides a StoreClient
// into an Fx application.
type FxStoreClientResult struct {
	fx.Out

	Client StoreClient
}

// NewFxStoreClient provides a StoreClient
// into an Fx application, using the given
// name for routing.
//
//  fx.Provide(
//    keyvaluepb.NewFxStoreClient("service-name"),
//    ...
//  )
// TODO(mensch): How will this work in v2?
func NewFxStoreClient(_ string, opts ...yarpcprotobuf.ClientOption) interface{} {
	return func(p FxStoreClientParams) FxStoreClientResult {
		return FxStoreClientResult{
			Client: NewStoreClient(p.Client, opts...),
		}
	}
}

// FxStoreServerParams defines the paramaters
// required to provide the StoreServer procedures
// into an Fx application.
type FxStoreServerParams struct {
	fx.In

	Server StoreServer
}

// FxStoreServerResult provides the StoreServer
// procedures into an Fx application.
type FxStoreServerResult struct {
	fx.Out

	Procedures []yarpc.Procedure `group:"yarpcfx"`
}

// NewFxStoreServer provides the StoreServer
// procedures to an Fx application. It expects
// a StoreServer to be present in the container.
//
//  fx.Provide(
//    keyvaluepb.NewFxStoreServer(),
//    ...
//  )
func NewFxStoreServer() interface{} {
	return func(p FxStoreServerParams) FxStoreServerResult {
		return FxStoreServerResult{
			Procedures: BuildStoreProcedures(p.Server),
		}
	}
}

func newStoreGetRequest() proto.Message  { return &common.GetRequest{} }
func newStoreGetResponse() proto.Message { return &common.GetResponse{} }
func newStoreSetRequest() proto.Message  { return &common.SetRequest{} }
func newStoreSetResponse() proto.Message { return &common.SetResponse{} }

var (
	_emptyStoreGetRequest  = &common.GetRequest{}
	_emptyStoreGetResponse = &common.GetResponse{}
	_emptyStoreSetRequest  = &common.SetRequest{}
	_emptyStoreSetResponse = &common.SetResponse{}
)

// Code generated by protoc-gen-yarpc-go
// source: example.proto
// DO NOT EDIT!

package examplepb

import (
	context "context"
	proto "github.com/gogo/protobuf/proto"
	fx "go.uber.org/fx"
	yarpc "go.uber.org/yarpc/v2"
	yarpcprotobuf "go.uber.org/yarpc/v2/yarpcprotobuf"
)

// ExampleClient is the Example service's client interface.
type ExampleClient interface {
	Foo(
		context.Context,
		*Request,
		...yarpc.CallOption,
	) (*Response, error)
	Bar(
		context.Context,
		...yarpc.CallOption,
	) (ExampleBarStreamClient, error)
	Baz(
		context.Context,
		*Request,
		...yarpc.CallOption,
	) (ExampleBazStreamClient, error)
	Qux(
		context.Context,
		...yarpc.CallOption,
	) (ExampleQuxStreamClient, error)
}

// NewExampleClient builds a new YARPC client for the Example service.
func NewExampleClient(c yarpc.Client, opts ...yarpcprotobuf.ClientOption) ExampleClient {
	return &_ExampleClient{stream: yarpcprotobuf.NewStreamClient(c, "example.Example", opts...)}
}

type _ExampleClient struct {
	stream yarpcprotobuf.StreamClient
}

var _ ExampleClient = (*_ExampleClient)(nil)

func (c *_ExampleClient) Foo(ctx context.Context, req *Request, opts ...yarpc.CallOption) (*Response, error) {
	msg, err := c.stream.Call(ctx, "Foo", req, newExampleFooResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*Response)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyExampleFooResponse, res)
	}
	return res, nil
}

func (c *_ExampleClient) Bar(ctx context.Context, opts ...yarpc.CallOption) (ExampleBarStreamClient, error) {
	s, err := c.stream.CallStream(ctx, "Bar", opts...)
	if err != nil {
		return err
	}
	return &_ExampleBarStreamClient{stream: s}, nil
}

func (c *_ExampleClient) Baz(ctx context.Context, req *Request, opts ...yarpc.CallOption) (ExampleBazStreamClient, error) {
	s, err := c.stream.CallStream(ctx, "Baz", opts...)
	if err != nil {
		return err
	}
	if err := s.Send(req); err != nil {
		return nil, err
	}
	return &_ExampleBazStreamClient{stream: s}, nil
}

func (c *_ExampleClient) Qux(ctx context.Context, opts ...yarpc.CallOption) (ExampleQuxStreamClient, error) {
	s, err := c.stream.CallStream(ctx, "Qux", opts...)
	if err != nil {
		return nil, err
	}
	return &_ExampleQuxStreamClient{stream: s}, nil
}

// ExampleBarStreamClient is a streaming interface used in the ExampleClient interface.
type ExampleBarStreamClient interface {
	Context() context.Context
	Send(*Request, ...yarpc.StreamOption) error
	CloseAndRecv(...yarpc.StreamOption) (*Response, error)
}

// ExampleBazStreamClient is a streaming interface used in the ExampleClient interface.
type ExampleBazStreamClient interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*Response, error)
	CloseSend(...yarpc.StreamOption) error
}

// ExampleQuxStreamClient is a streaming interface used in the ExampleClient interface.
type ExampleQuxStreamClient interface {
	Context() context.Context
	Send(*Request, ...yarpc.StreamOption) error
	Recv(...yarpc.StreamOption) (*Response, error)
	CloseSend(...yarpc.StreamOption) error
}

type _ExampleBarStreamClient struct {
	stream *yarpcprotobuf.StreamClient
}

var _ ExampleBarStreamClient = (*_ExampleBarStreamClient)(nil)

func (c *_ExampleBarStreamClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_ExampleBarStreamClient) Send(req *Request, opts ...yarpc.StreamOption) error {
	return c.stream.Send(req, opts...)
}

func (c *_ExampleBarStreamClient) CloseAndRecv(opts ...yarpc.StreamOption) (*Request, error) {
	if err := c.stream.Close(opts...); err != nil {
		return nil, err
	}
	msg, err := c.stream.Receive(newExampleBarResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*Response)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyExampleBarResponse, msg)
	}
	return res, err
}

type _ExampleBazStreamClient struct {
	stream *yarpcprotobuf.StreamClient
}

var _ ExampleBazStreamClient = (*_ExampleBazStreamClient)(nil)

func (c *_ExampleBazStreamClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_ExampleBazStreamClient) Recv(opts ...yarpc.StreamOption) (*Response, error) {
	msg, err := c.stream.Receive(newExampleBazResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*Response)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyExampleBazResponse, msg)
	}
	return res, nil
}

func (c *_ExampleBazStreamClient) CloseSend(opts ...yarpc.StreamOption) error {
	return c.stream.Close(opts...)
}

type _ExampleQuxStreamClient struct {
	stream *yarpcprotobuf.StreamClient
}

var _ ExampleQuxStreamClient = (*_ExampleQuxStreamClient)(nil)

func (c *_ExampleQuxStreamClient) Context() context.Context {
	return c.stream.Context()
}

func (c *_ExampleQuxStreamClient) Send(req *Request, opts ...yarpc.StreamOption) error {
	return c.stream.Send(req, opts...)
}

func (c *_ExampleQuxStreamClient) Recv(opts ...yarpc.StreamOption) (*Response, error) {
	msg, err := c.stream.Receive(newExampleQuxResponse, opts...)
	if err != nil {
		return nil, err
	}
	res, ok := msg.(*Response)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyExampleQuxResponse, msg)
	}
	return res, nil
}

func (c *_ExampleQuxStreamClient) CloseSend(opts ...yarpc.StreamOption) error {
	return c.stream.Close(opts...)
}

// ExampleServer is the Example service's server interface.
type ExampleServer interface {
	Foo(
		context.Context,
		*Request,
	) (*Response, error)
	Bar(
		ExampleBarStreamServer,
	) (ExampleBarStreamServer, error)
	Baz(
		*Request,
		ExampleBazStreamServer,
	) error
	Qux(
		ExampleQuxStreamServer,
	) error
}

// BuildExampleProcedures constructs the YARPC procedures for the Example service.
func BuildExampleProcedures(s ExampleServer) []yarpc.Procedure {
	h := &_ExampleServer{server: s}
	return yarpcprotobuf.Procedures(
		yarpcprotobuf.ProceduresParams{
			Service: "example.Example",
			Unary: []yarpcprotobuf.UnaryProceduresParams{
				{
					MethodName: "Foo",
					Handler: yarpcprotobuf.NewUnaryHandler{
						yarpcprotobuf.UnaryHandlerParams{
							Handle: h.Foo,
							Create: newExampleFooRequest(),
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

type _ExampleServer struct {
	server ExampleServer
}

var _ ExampleServer = (*_ExampleServer)(nil)

func (h *_ExampleServer) Foo(ctx context.Context, m proto.Message) (proto.Message, error) {
	req, _ := m.(*Request)
	if req == nil {
		return nil, yarpcprotobuf.CastError(_emptyExampleFooRequest, m)
	}
	return h.server.Foo(ctx, req)
}

func (h *_ExampleServer) Bar(s *yarpcprotobuf.StreamServer) error {
	res, err := h.server.Bar(&_ExampleBarStreamServer{server: s})
	if err != nil {
		return err
	}
	return s.Send(res)
}

func (h *_ExampleServer) Baz(s *yarpcprotobuf.StreamServer) error {
	recv, err := s.Receive(newExampleBazRequest)
	if err != nil {
		return err
	}
	req, _ := recv.(*Request)
	if req == nil {
		return yarpcprotobuf.CastError(_emptyExampleBazRequest, recv)
	}
	return h.server.Baz(req, &_ExampleBazStreamServer{server: s})
}

func (h *_ExampleServer) Qux(s *yarpcprotobuf.StreamServer) error {
	return h.server.Qux(&_ExampleQuxStreamServer{stream: s})
}

// ExampleBarStreamServer is a streaming interface used in the ExampleServer interface.
type ExampleBarStreamServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*Request, error)
}

// ExampleBazStreamServer is a streaming interface used in the ExampleServer interface.
type ExampleBazStreamServer interface {
	Context() context.Context
	Send(*Response, ...yarpc.StreamOption) error
}

// ExampleQuxStreamServer is a streaming interface used in the ExampleServer interface.
type ExampleQuxStreamServer interface {
	Context() context.Context
	Recv(...yarpc.StreamOption) (*Request, error)
	Send(*Response, ...yarpc.StreamOption) error
}

type _ExampleBarStreamServer struct {
	stream *yarpcprotobuf.StreamServer
}

var _ ExampleBarStreamServer = (*_ExampleBarStreamServer)(nil)

func (s *_ExampleBarStreamServer) Context() context.Context {
	return s.stream.Context()
}

func (s *_ExampleBarStreamServer) Recv(opts ...yarpc.StreamOption) (*Request, error) {
	msg, err := s.stream.Receive(newExampleBarRequest, opts...)
	if err != nil {
		return nil, err
	}
	req, ok := msg.(*Request)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyExampleBarRequest, msg)
	}
	return req, nil
}

type _ExampleBazStreamServer struct {
	stream *yarpcprotobuf.StreamServer
}

var _ ExampleBazStreamServer = (*_ExampleBazStreamServer)(nil)

func (s *_ExampleBazStreamServer) Context() context.Context {
	return s.stream.Context()
}

func (s *_ExampleBazStreamServer) Send(res *Response, opts ...yarpc.StreamOption) error {
	return s.stream.Send(res, opts...)
}

type _ExampleQuxStreamServer struct {
	stream *yarpcprotobuf.StreamServer
}

var _ ExampleQuxStreamServer = (*_ExampleQuxStreamServer)(nil)

func (s *_ExampleQuxStreamServer) Context() context.Context {
	return s.stream.Context()
}

func (s *_ExampleQuxStreamServer) Recv(opts ...yarpc.StreamOption) (*Request, error) {
	msg, err := s.stream.Receive(newExampleQuxRequest, opts...)
	if err != nil {
		return nil, err
	}
	req, ok := msg.(*Request)
	if !ok {
		return nil, yarpcprotobuf.CastError(_emptyExampleQuxRequest, msg)
	}
	return req, nil
}

func (s *_ExampleQuxStreamServer) Send(res *Response, opts ...yarpc.StreamOption) error {
	return s.stream.Send(res, opts...)
}

// FxExampleClientParams defines the parameters
// required to provide a ExampleClient into an
// Fx application.
type FxExampleClientParams struct {
	fx.In

	Client yarpc.Client
}

// FxExampleClientResult provides a ExampleClient
// into an Fx application.
type FxExampleClientResult struct {
	fx.Out

	Client ExampleClient
}

// NewFxExampleClient provides a ExampleClient
// into an Fx application, using the given
// name for routing.
//
//  fx.Provide(
//    examplepb.NewFxExampleClient("service-name"),
//    ...
//  )
// TODO(mensch): How will this work in v2?
func NewFxExampleClient(_ string, opts ...yarpcprotobuf.ClientOption) interface{} {
	return func(p FxExampleClientParams) FxExampleClientResult {
		return FxExampleClientResult{
			Client: NewFxExampleClient(p.Client, opts...),
		}
	}
}

// FxExampleServerParams defines the paramaters
// required to provide the ExampleServer procedures
// into an Fx application.
type FxExampleServerParams struct {
	fx.In

	Server ExampleServer
}

// FxExampleServerResult provides the ExampleServer
// procedures into an Fx application.
type FxExampleServerResult struct {
	fx.Out

	Procedures []yarpc.Procedure `group:"yarpcfx"`
}

// NewFxExampleServer provides the ExampleServer
// procedures to an Fx application. It expects
// a ExampleServer to be present in the container.
//
//  fx.Provide(
//    examplepb.NewFxExampleServer(),
//    ...
//  )
func NewFxExampleServer() interface{} {
	return func(p FxExampleServerParams) FxExampleServerResult {
		return FxExampleServerResult{
			Procedures: BuildExampleProcedures(p.Server),
		}
	}
}

func newExampleFooRequest()  { return &Request{} }
func newExampleFooResponse() { return &Response{} }
func newExampleBarRequest()  { return &Request{} }
func newExampleBarResponse() { return &Response{} }
func newExampleBazRequest()  { return &Request{} }
func newExampleBazResponse() { return &Response{} }
func newExampleQuxRequest()  { return &Request{} }
func newExampleQuxResponse() { return &Response{} }

var (
	_emptyExampleFooRequest  = &Request{}
	_emptyExampleFooResponse = &Response{}
	_emptyExampleBarRequest  = &Request{}
	_emptyExampleBarResponse = &Response{}
	_emptyExampleBazRequest  = &Request{}
	_emptyExampleBazResponse = &Response{}
	_emptyExampleQuxRequest  = &Request{}
	_emptyExampleQuxResponse = &Response{}
)

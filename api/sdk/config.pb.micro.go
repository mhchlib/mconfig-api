// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/v0/config.proto

package sdk

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for MConfig service

func NewMConfigEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MConfig service

type MConfigService interface {
	GetVStream(ctx context.Context, in *GetVRequest, opts ...client.CallOption) (MConfig_GetVStreamService, error)
}

type mConfigService struct {
	c    client.Client
	name string
}

func NewMConfigService(name string, c client.Client) MConfigService {
	return &mConfigService{
		c:    c,
		name: name,
	}
}

func (c *mConfigService) GetVStream(ctx context.Context, in *GetVRequest, opts ...client.CallOption) (MConfig_GetVStreamService, error) {
	req := c.c.NewRequest(c.name, "MConfig.GetVStream", &GetVRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &mConfigServiceGetVStream{stream}, nil
}

type MConfig_GetVStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*GetVResponse, error)
}

type mConfigServiceGetVStream struct {
	stream client.Stream
}

func (x *mConfigServiceGetVStream) Close() error {
	return x.stream.Close()
}

func (x *mConfigServiceGetVStream) Context() context.Context {
	return x.stream.Context()
}

func (x *mConfigServiceGetVStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *mConfigServiceGetVStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *mConfigServiceGetVStream) Recv() (*GetVResponse, error) {
	m := new(GetVResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for MConfig service

type MConfigHandler interface {
	GetVStream(context.Context, *GetVRequest, MConfig_GetVStreamStream) error
}

func RegisterMConfigHandler(s server.Server, hdlr MConfigHandler, opts ...server.HandlerOption) error {
	type mConfig interface {
		GetVStream(ctx context.Context, stream server.Stream) error
	}
	type MConfig struct {
		mConfig
	}
	h := &mConfigHandler{hdlr}
	return s.Handle(s.NewHandler(&MConfig{h}, opts...))
}

type mConfigHandler struct {
	MConfigHandler
}

func (h *mConfigHandler) GetVStream(ctx context.Context, stream server.Stream) error {
	m := new(GetVRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.MConfigHandler.GetVStream(ctx, m, &mConfigGetVStreamStream{stream})
}

type MConfig_GetVStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*GetVResponse) error
}

type mConfigGetVStreamStream struct {
	stream server.Stream
}

func (x *mConfigGetVStreamStream) Close() error {
	return x.stream.Close()
}

func (x *mConfigGetVStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *mConfigGetVStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *mConfigGetVStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *mConfigGetVStreamStream) Send(m *GetVResponse) error {
	return x.stream.Send(m)
}

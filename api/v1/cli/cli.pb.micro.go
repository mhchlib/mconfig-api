// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/v1/cli.proto

package cli

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mhchlib/mconfig-api/api/v1/common"
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

// Api Endpoints for MConfigCli service

func NewMConfigCliEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for MConfigCli service

type MConfigCliService interface {
	PutMconfigConfig(ctx context.Context, in *PutMconfigRequest, opts ...client.CallOption) (*PutMconfigResponse, error)
	DeleteMconfigConfig(ctx context.Context, in *DeleteMconfigConfigRequest, opts ...client.CallOption) (*DeleteMconfigConfigResponse, error)
	InitMconfigApp(ctx context.Context, in *InitMconfigAppRequest, opts ...client.CallOption) (*InitMconfigAppResponse, error)
	UpdateMconfigApp(ctx context.Context, in *UpdateMconfigAppRequest, opts ...client.CallOption) (*UpdateMconfigAppResponse, error)
	DeleteMconfigApp(ctx context.Context, in *DeleteMconfigAppRequest, opts ...client.CallOption) (*DeleteMconfigAppResponse, error)
}

type mConfigCliService struct {
	c    client.Client
	name string
}

func NewMConfigCliService(name string, c client.Client) MConfigCliService {
	return &mConfigCliService{
		c:    c,
		name: name,
	}
}

func (c *mConfigCliService) PutMconfigConfig(ctx context.Context, in *PutMconfigRequest, opts ...client.CallOption) (*PutMconfigResponse, error) {
	req := c.c.NewRequest(c.name, "MConfigCli.PutMconfigConfig", in)
	out := new(PutMconfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mConfigCliService) DeleteMconfigConfig(ctx context.Context, in *DeleteMconfigConfigRequest, opts ...client.CallOption) (*DeleteMconfigConfigResponse, error) {
	req := c.c.NewRequest(c.name, "MConfigCli.DeleteMconfigConfig", in)
	out := new(DeleteMconfigConfigResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mConfigCliService) InitMconfigApp(ctx context.Context, in *InitMconfigAppRequest, opts ...client.CallOption) (*InitMconfigAppResponse, error) {
	req := c.c.NewRequest(c.name, "MConfigCli.InitMconfigApp", in)
	out := new(InitMconfigAppResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mConfigCliService) UpdateMconfigApp(ctx context.Context, in *UpdateMconfigAppRequest, opts ...client.CallOption) (*UpdateMconfigAppResponse, error) {
	req := c.c.NewRequest(c.name, "MConfigCli.UpdateMconfigApp", in)
	out := new(UpdateMconfigAppResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mConfigCliService) DeleteMconfigApp(ctx context.Context, in *DeleteMconfigAppRequest, opts ...client.CallOption) (*DeleteMconfigAppResponse, error) {
	req := c.c.NewRequest(c.name, "MConfigCli.DeleteMconfigApp", in)
	out := new(DeleteMconfigAppResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MConfigCli service

type MConfigCliHandler interface {
	PutMconfigConfig(context.Context, *PutMconfigRequest, *PutMconfigResponse) error
	DeleteMconfigConfig(context.Context, *DeleteMconfigConfigRequest, *DeleteMconfigConfigResponse) error
	InitMconfigApp(context.Context, *InitMconfigAppRequest, *InitMconfigAppResponse) error
	UpdateMconfigApp(context.Context, *UpdateMconfigAppRequest, *UpdateMconfigAppResponse) error
	DeleteMconfigApp(context.Context, *DeleteMconfigAppRequest, *DeleteMconfigAppResponse) error
}

func RegisterMConfigCliHandler(s server.Server, hdlr MConfigCliHandler, opts ...server.HandlerOption) error {
	type mConfigCli interface {
		PutMconfigConfig(ctx context.Context, in *PutMconfigRequest, out *PutMconfigResponse) error
		DeleteMconfigConfig(ctx context.Context, in *DeleteMconfigConfigRequest, out *DeleteMconfigConfigResponse) error
		InitMconfigApp(ctx context.Context, in *InitMconfigAppRequest, out *InitMconfigAppResponse) error
		UpdateMconfigApp(ctx context.Context, in *UpdateMconfigAppRequest, out *UpdateMconfigAppResponse) error
		DeleteMconfigApp(ctx context.Context, in *DeleteMconfigAppRequest, out *DeleteMconfigAppResponse) error
	}
	type MConfigCli struct {
		mConfigCli
	}
	h := &mConfigCliHandler{hdlr}
	return s.Handle(s.NewHandler(&MConfigCli{h}, opts...))
}

type mConfigCliHandler struct {
	MConfigCliHandler
}

func (h *mConfigCliHandler) PutMconfigConfig(ctx context.Context, in *PutMconfigRequest, out *PutMconfigResponse) error {
	return h.MConfigCliHandler.PutMconfigConfig(ctx, in, out)
}

func (h *mConfigCliHandler) DeleteMconfigConfig(ctx context.Context, in *DeleteMconfigConfigRequest, out *DeleteMconfigConfigResponse) error {
	return h.MConfigCliHandler.DeleteMconfigConfig(ctx, in, out)
}

func (h *mConfigCliHandler) InitMconfigApp(ctx context.Context, in *InitMconfigAppRequest, out *InitMconfigAppResponse) error {
	return h.MConfigCliHandler.InitMconfigApp(ctx, in, out)
}

func (h *mConfigCliHandler) UpdateMconfigApp(ctx context.Context, in *UpdateMconfigAppRequest, out *UpdateMconfigAppResponse) error {
	return h.MConfigCliHandler.UpdateMconfigApp(ctx, in, out)
}

func (h *mConfigCliHandler) DeleteMconfigApp(ctx context.Context, in *DeleteMconfigAppRequest, out *DeleteMconfigAppResponse) error {
	return h.MConfigCliHandler.DeleteMconfigApp(ctx, in, out)
}

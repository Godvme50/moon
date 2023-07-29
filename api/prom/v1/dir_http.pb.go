// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.19.4
// source: prom/v1/dir.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDirCreateDir = "/api.prom.Dir/CreateDir"
const OperationDirDeleteDir = "/api.prom.Dir/DeleteDir"
const OperationDirGetDir = "/api.prom.Dir/GetDir"
const OperationDirListDir = "/api.prom.Dir/ListDir"
const OperationDirUpdateDir = "/api.prom.Dir/UpdateDir"

type DirHTTPServer interface {
	CreateDir(context.Context, *CreateDirRequest) (*CreateDirReply, error)
	DeleteDir(context.Context, *DeleteDirRequest) (*DeleteDirReply, error)
	GetDir(context.Context, *GetDirRequest) (*GetDirReply, error)
	ListDir(context.Context, *ListDirRequest) (*ListDirReply, error)
	UpdateDir(context.Context, *UpdateDirRequest) (*UpdateDirReply, error)
}

func RegisterDirHTTPServer(s *http.Server, srv DirHTTPServer) {
	r := s.Route("/")
	r.POST("/prom/v1/dir/add", _Dir_CreateDir0_HTTP_Handler(srv))
	r.PUT("/prom/v1/dir/edit/{id}", _Dir_UpdateDir0_HTTP_Handler(srv))
	r.DELETE("/prom/v1/dir/delete/{id}", _Dir_DeleteDir0_HTTP_Handler(srv))
	r.GET("/prom/v1/dir/{id}", _Dir_GetDir0_HTTP_Handler(srv))
	r.POST("/prom/v1/dir/list", _Dir_ListDir0_HTTP_Handler(srv))
}

func _Dir_CreateDir0_HTTP_Handler(srv DirHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDirRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDirCreateDir)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDir(ctx, req.(*CreateDirRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDirReply)
		return ctx.Result(200, reply)
	}
}

func _Dir_UpdateDir0_HTTP_Handler(srv DirHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDirRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDirUpdateDir)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDir(ctx, req.(*UpdateDirRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDirReply)
		return ctx.Result(200, reply)
	}
}

func _Dir_DeleteDir0_HTTP_Handler(srv DirHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDirRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDirDeleteDir)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDir(ctx, req.(*DeleteDirRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDirReply)
		return ctx.Result(200, reply)
	}
}

func _Dir_GetDir0_HTTP_Handler(srv DirHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDirRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDirGetDir)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDir(ctx, req.(*GetDirRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDirReply)
		return ctx.Result(200, reply)
	}
}

func _Dir_ListDir0_HTTP_Handler(srv DirHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDirRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDirListDir)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDir(ctx, req.(*ListDirRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDirReply)
		return ctx.Result(200, reply)
	}
}

type DirHTTPClient interface {
	CreateDir(ctx context.Context, req *CreateDirRequest, opts ...http.CallOption) (rsp *CreateDirReply, err error)
	DeleteDir(ctx context.Context, req *DeleteDirRequest, opts ...http.CallOption) (rsp *DeleteDirReply, err error)
	GetDir(ctx context.Context, req *GetDirRequest, opts ...http.CallOption) (rsp *GetDirReply, err error)
	ListDir(ctx context.Context, req *ListDirRequest, opts ...http.CallOption) (rsp *ListDirReply, err error)
	UpdateDir(ctx context.Context, req *UpdateDirRequest, opts ...http.CallOption) (rsp *UpdateDirReply, err error)
}

type DirHTTPClientImpl struct {
	cc *http.Client
}

func NewDirHTTPClient(client *http.Client) DirHTTPClient {
	return &DirHTTPClientImpl{client}
}

func (c *DirHTTPClientImpl) CreateDir(ctx context.Context, in *CreateDirRequest, opts ...http.CallOption) (*CreateDirReply, error) {
	var out CreateDirReply
	pattern := "/prom/v1/dir/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDirCreateDir))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DirHTTPClientImpl) DeleteDir(ctx context.Context, in *DeleteDirRequest, opts ...http.CallOption) (*DeleteDirReply, error) {
	var out DeleteDirReply
	pattern := "/prom/v1/dir/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDirDeleteDir))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DirHTTPClientImpl) GetDir(ctx context.Context, in *GetDirRequest, opts ...http.CallOption) (*GetDirReply, error) {
	var out GetDirReply
	pattern := "/prom/v1/dir/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDirGetDir))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DirHTTPClientImpl) ListDir(ctx context.Context, in *ListDirRequest, opts ...http.CallOption) (*ListDirReply, error) {
	var out ListDirReply
	pattern := "/prom/v1/dir/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDirListDir))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DirHTTPClientImpl) UpdateDir(ctx context.Context, in *UpdateDirRequest, opts ...http.CallOption) (*UpdateDirReply, error) {
	var out UpdateDirReply
	pattern := "/prom/v1/dir/edit/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDirUpdateDir))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

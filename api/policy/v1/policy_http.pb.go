// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.0

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

type PolicyHTTPServer interface {
	CreatePolicy(context.Context, *PolicyStruct) (*PolicyStruct, error)
	DeletePolicy(context.Context, *DeletePolicyRequest) (*PolicyStruct, error)
	GetPolicy(context.Context, *GetPolicyRequest) (*PolicyStruct, error)
	ListPolicy(context.Context, *ListPolicyRequest) (*ListPolicyReply, error)
	UpdatePolicy(context.Context, *PolicyStruct) (*PolicyStruct, error)
}

func RegisterPolicyHTTPServer(s *http.Server, srv PolicyHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/basePolicy", _Policy_CreatePolicy0_HTTP_Handler(srv))
	r.GET("/v1/basePolicy/{id}", _Policy_GetPolicy1_HTTP_Handler(srv))
	r.GET("/v1/basePolicies", _Policy_ListPolicy0_HTTP_Handler(srv))
	r.PUT("/v1/basePolicy/{id}", _Policy_UpdatePolicy0_HTTP_Handler(srv))
	r.DELETE("/v1/basePolicy/{id}", _Policy_DeletePolicy0_HTTP_Handler(srv))
}

func _Policy_CreatePolicy0_HTTP_Handler(srv PolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PolicyStruct
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.Policy/CreatePolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePolicy(ctx, req.(*PolicyStruct))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PolicyStruct)
		return ctx.Result(200, reply)
	}
}

func _Policy_GetPolicy1_HTTP_Handler(srv PolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetPolicyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.Policy/GetPolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPolicy(ctx, req.(*GetPolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PolicyStruct)
		return ctx.Result(200, reply)
	}
}

func _Policy_ListPolicy0_HTTP_Handler(srv PolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListPolicyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.Policy/ListPolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPolicy(ctx, req.(*ListPolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListPolicyReply)
		return ctx.Result(200, reply)
	}
}

func _Policy_UpdatePolicy0_HTTP_Handler(srv PolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PolicyStruct
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.Policy/UpdatePolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePolicy(ctx, req.(*PolicyStruct))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PolicyStruct)
		return ctx.Result(200, reply)
	}
}

func _Policy_DeletePolicy0_HTTP_Handler(srv PolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeletePolicyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.Policy/DeletePolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePolicy(ctx, req.(*DeletePolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PolicyStruct)
		return ctx.Result(200, reply)
	}
}

type PolicyHTTPClient interface {
	CreatePolicy(ctx context.Context, req *PolicyStruct, opts ...http.CallOption) (rsp *PolicyStruct, err error)
	DeletePolicy(ctx context.Context, req *DeletePolicyRequest, opts ...http.CallOption) (rsp *PolicyStruct, err error)
	GetPolicy(ctx context.Context, req *GetPolicyRequest, opts ...http.CallOption) (rsp *PolicyStruct, err error)
	ListPolicy(ctx context.Context, req *ListPolicyRequest, opts ...http.CallOption) (rsp *ListPolicyReply, err error)
	UpdatePolicy(ctx context.Context, req *PolicyStruct, opts ...http.CallOption) (rsp *PolicyStruct, err error)
}

type PolicyHTTPClientImpl struct {
	cc *http.Client
}

func NewPolicyHTTPClient(client *http.Client) PolicyHTTPClient {
	return &PolicyHTTPClientImpl{client}
}

func (c *PolicyHTTPClientImpl) CreatePolicy(ctx context.Context, in *PolicyStruct, opts ...http.CallOption) (*PolicyStruct, error) {
	var out PolicyStruct
	pattern := "/v1/basePolicy"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.policy.v1.Policy/CreatePolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PolicyHTTPClientImpl) DeletePolicy(ctx context.Context, in *DeletePolicyRequest, opts ...http.CallOption) (*PolicyStruct, error) {
	var out PolicyStruct
	pattern := "/v1/basePolicy/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.policy.v1.Policy/DeletePolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PolicyHTTPClientImpl) GetPolicy(ctx context.Context, in *GetPolicyRequest, opts ...http.CallOption) (*PolicyStruct, error) {
	var out PolicyStruct
	pattern := "/v1/basePolicy/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.policy.v1.Policy/GetPolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PolicyHTTPClientImpl) ListPolicy(ctx context.Context, in *ListPolicyRequest, opts ...http.CallOption) (*ListPolicyReply, error) {
	var out ListPolicyReply
	pattern := "/v1/basePolicies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.policy.v1.Policy/ListPolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PolicyHTTPClientImpl) UpdatePolicy(ctx context.Context, in *PolicyStruct, opts ...http.CallOption) (*PolicyStruct, error) {
	var out PolicyStruct
	pattern := "/v1/basePolicy/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.policy.v1.Policy/UpdatePolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

type ImportPolicyHTTPServer interface {
	CreatePolicy(context.Context, *ImportPolicyStruct) (*ImportPolicyStruct, error)
	DeletePolicy(context.Context, *DeleteImportPolicyRequest) (*ImportPolicyStruct, error)
	GetPolicy(context.Context, *GetImportPolicyRequest) (*ImportPolicyStruct, error)
	ListPolicy(context.Context, *ListImportPolicyRequest) (*ListImportPolicyReply, error)
	UpdatePolicy(context.Context, *ImportPolicyStruct) (*ImportPolicyStruct, error)
}

func RegisterImportPolicyHTTPServer(s *http.Server, srv ImportPolicyHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/importPolicy", _ImportPolicy_CreatePolicy1_HTTP_Handler(srv))
	r.GET("/v1/importPolicy/{id}", _ImportPolicy_GetPolicy2_HTTP_Handler(srv))
	r.GET("/v1/importPolicies", _ImportPolicy_ListPolicy1_HTTP_Handler(srv))
	r.PUT("/v1/importPolicy/{id}", _ImportPolicy_UpdatePolicy1_HTTP_Handler(srv))
	r.DELETE("/v1/importPolicy/{id}", _ImportPolicy_DeletePolicy1_HTTP_Handler(srv))
}

func _ImportPolicy_CreatePolicy1_HTTP_Handler(srv ImportPolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ImportPolicyStruct
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.ImportPolicy/CreatePolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePolicy(ctx, req.(*ImportPolicyStruct))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ImportPolicyStruct)
		return ctx.Result(200, reply)
	}
}

func _ImportPolicy_GetPolicy2_HTTP_Handler(srv ImportPolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetImportPolicyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.ImportPolicy/GetPolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetPolicy(ctx, req.(*GetImportPolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ImportPolicyStruct)
		return ctx.Result(200, reply)
	}
}

func _ImportPolicy_ListPolicy1_HTTP_Handler(srv ImportPolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListImportPolicyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.ImportPolicy/ListPolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListPolicy(ctx, req.(*ListImportPolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListImportPolicyReply)
		return ctx.Result(200, reply)
	}
}

func _ImportPolicy_UpdatePolicy1_HTTP_Handler(srv ImportPolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ImportPolicyStruct
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.ImportPolicy/UpdatePolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdatePolicy(ctx, req.(*ImportPolicyStruct))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ImportPolicyStruct)
		return ctx.Result(200, reply)
	}
}

func _ImportPolicy_DeletePolicy1_HTTP_Handler(srv ImportPolicyHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteImportPolicyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.policy.v1.ImportPolicy/DeletePolicy")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeletePolicy(ctx, req.(*DeleteImportPolicyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ImportPolicyStruct)
		return ctx.Result(200, reply)
	}
}

type ImportPolicyHTTPClient interface {
	CreatePolicy(ctx context.Context, req *ImportPolicyStruct, opts ...http.CallOption) (rsp *ImportPolicyStruct, err error)
	DeletePolicy(ctx context.Context, req *DeleteImportPolicyRequest, opts ...http.CallOption) (rsp *ImportPolicyStruct, err error)
	GetPolicy(ctx context.Context, req *GetImportPolicyRequest, opts ...http.CallOption) (rsp *ImportPolicyStruct, err error)
	ListPolicy(ctx context.Context, req *ListImportPolicyRequest, opts ...http.CallOption) (rsp *ListImportPolicyReply, err error)
	UpdatePolicy(ctx context.Context, req *ImportPolicyStruct, opts ...http.CallOption) (rsp *ImportPolicyStruct, err error)
}

type ImportPolicyHTTPClientImpl struct {
	cc *http.Client
}

func NewImportPolicyHTTPClient(client *http.Client) ImportPolicyHTTPClient {
	return &ImportPolicyHTTPClientImpl{client}
}

func (c *ImportPolicyHTTPClientImpl) CreatePolicy(ctx context.Context, in *ImportPolicyStruct, opts ...http.CallOption) (*ImportPolicyStruct, error) {
	var out ImportPolicyStruct
	pattern := "/v1/importPolicy"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.policy.v1.ImportPolicy/CreatePolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ImportPolicyHTTPClientImpl) DeletePolicy(ctx context.Context, in *DeleteImportPolicyRequest, opts ...http.CallOption) (*ImportPolicyStruct, error) {
	var out ImportPolicyStruct
	pattern := "/v1/importPolicy/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.policy.v1.ImportPolicy/DeletePolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ImportPolicyHTTPClientImpl) GetPolicy(ctx context.Context, in *GetImportPolicyRequest, opts ...http.CallOption) (*ImportPolicyStruct, error) {
	var out ImportPolicyStruct
	pattern := "/v1/importPolicy/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.policy.v1.ImportPolicy/GetPolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ImportPolicyHTTPClientImpl) ListPolicy(ctx context.Context, in *ListImportPolicyRequest, opts ...http.CallOption) (*ListImportPolicyReply, error) {
	var out ListImportPolicyReply
	pattern := "/v1/importPolicies"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.policy.v1.ImportPolicy/ListPolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ImportPolicyHTTPClientImpl) UpdatePolicy(ctx context.Context, in *ImportPolicyStruct, opts ...http.CallOption) (*ImportPolicyStruct, error) {
	var out ImportPolicyStruct
	pattern := "/v1/importPolicy/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.policy.v1.ImportPolicy/UpdatePolicy"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

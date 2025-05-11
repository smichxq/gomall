package service

import (
	"context"

	category "github.com/cloudwego/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/gomall/app/frontend/infra/rpc"
	rpcProduct "github.com/cloudwego/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	p, err := rpc.ProductClient.ListProducts(h.Context, &rpcProduct.ListProductsReq{
		CategroyName: req.Category,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}

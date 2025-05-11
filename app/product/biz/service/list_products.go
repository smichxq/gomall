package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/product/biz/model"
	product "github.com/cloudwego/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.

	if req.CategroyName == "" {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "categroy name is required")
	}

	categories, err := model.NewCategoryQuery(s.ctx, mysql.DB).GetProductsByCategoryName(req.GetCategroyName())
	if err != nil {
		return nil, err
	}

	respp := &product.ListProductsResp{}

	for _, v := range categories {
		for _, p := range v.Products {
			respp.Products = append(respp.Products, &product.Product{
				Id:          uint32(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Price:       fmt.Sprintf("%.2f", p.Price),
			})
		}
	}

	return respp, nil
}

package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/gomall/app/product/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/product/biz/model"
	product "github.com/cloudwego/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req == nil {
		return nil, kerrors.NewGRPCBizStatusError(2004001, "product id is required")
	}

	query := model.NewProductQuery(s.ctx, mysql.DB)

	res, err := query.QeryById(req.Id)
	if err != nil {
		return nil, err
	}

	// Helper function to map []model.Category to []string
	mapCategoriesToStrings := func(categories []model.Category) []string {
		strings := make([]string, len(categories))
		for i, category := range categories {
			strings[i] = category.Name // Assuming model.Category has a Name field
		}
		return strings
	}

	categories := mapCategoriesToStrings(res.Categories)

	return &product.GetProductResp{
		Product: &product.Product{
			Id:         uint32(res.ID),
			Name:       res.Name,
			Categories: categories,
			Picture:    res.Picture,
			Price:      fmt.Sprintf("%.2f", res.Price),
		},
	}, nil
}

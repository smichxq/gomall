package service

import (
	"context"

	"github.com/cloudwego/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/cart/biz/model"
	cart "github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.

	list, err := model.GetCartByUserId(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50001, err.Error())
	}

	var items []*cart.CartItem
	for _, item := range list {
		items = append(items, &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Qty,
		})
	}

	return &cart.GetCartResp{Items: items}, nil
}

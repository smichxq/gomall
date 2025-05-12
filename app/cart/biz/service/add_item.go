package service

import (
	"context"

	"github.com/cloudwego/gomall/app/cart/biz/dal/mysql"
	"github.com/cloudwego/gomall/app/cart/biz/model"
	"github.com/cloudwego/gomall/app/cart/infra/rpc"
	cart "github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.

	// 调用product服务获取商品信息
	respp, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}

	if respp == nil || respp.Product.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(40040, "product not found")
	}

	err = model.AddItem(s.ctx, mysql.DB, &model.Cart{
		UserId:    req.UserId,
		ProductId: respp.Product.Id,
		Qty:       req.Item.Quantity,
	})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(50000, "add cart fail")
	}

	return &cart.AddItemResp{}, nil
}

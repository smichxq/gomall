package service

import (
	"context"
	"strconv"

	common "github.com/cloudwego/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/gomall/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/gomall/app/frontend/utils"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context))})
	if err != nil {
		frontendUtils.MustHandleErr(err)
		return nil, kerrors.NewGRPCBizStatusError(400020, "Cart client invoke fail")
	}

	var items []map[string]string
	var total float64
	for _, item := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: item.ProductId})
		if err != nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Price":       strconv.FormatFloat(func() float64 { price, _ := strconv.ParseFloat(p.Price, 64); return price }(), 'f', 2, 64),
			"Picture":     p.Picture,
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		price, err := strconv.ParseFloat(p.Price, 64)
		if err != nil {
			continue
		}
		total += price * float64(item.Quantity)

	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}

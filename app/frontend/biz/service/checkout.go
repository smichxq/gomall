package service

import (
	"context"
	"strconv"

	common "github.com/cloudwego/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/gomall/app/frontend/infra/rpc"
	frontendutils "github.com/cloudwego/gomall/app/frontend/utils"
	rpccart "github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/cloudwego/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	var items []map[string]string
	userId := frontendutils.GetUserIdFromCtx(h.Context)

	carts, err := rpc.CartClient.GetCart(h.Context, &rpccart.GetCartReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}
	var total float32

	for _, v := range carts.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{
			Id: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name": p.Name,
			"Price": func() string {
				price, _ := strconv.ParseFloat(p.Price, 64)
				return strconv.FormatFloat(price, 'f', 2, 64)
			}(),
			"Picture": p.Picture,
			"Qty":     strconv.Itoa(int(v.Quantity)),
		})
		price, err := strconv.ParseFloat(p.Price, 32)
		if err != nil {
			return nil, err
		}
		total += float32(v.Quantity) * float32(price)
	}

	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}

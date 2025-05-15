package service

import (
	"context"
	"strconv"

	"github.com/cloudwego/gomall/app/checkout/infra/mq"
	"github.com/cloudwego/gomall/app/checkout/infra/rpc"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/cart"
	checkout "github.com/cloudwego/gomall/rpc_gen/kitex_gen/checkout"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/email"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/gomall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.

	// 获取用户购物车
	cartResult, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{UserId: req.UserId})
	if err != nil {
		return nil, kerrors.NewGRPCBizStatusError(5005001, err.Error())
	}
	if cartResult == nil || cartResult.Items == nil {
		return nil, kerrors.NewGRPCBizStatusError(5004001, "cart is empty")
	}

	// 总额
	var total float32

	// 循环获取商品详情
	for _, cartItem := range cartResult.Items {
		productResp, resultErr := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: cartItem.ProductId,
		})

		if resultErr != nil {
			return nil, resultErr
		}

		if productResp.Product == nil {
			continue
		}

		p, err := strconv.ParseFloat(productResp.Product.Price, 32)
		if err != nil {
			return nil, kerrors.NewGRPCBizStatusError(5004002, "invalid product price")
		}

		// 计算金额
		cost := float32(p) * float32(cartItem.Quantity)
		total += cost
	}

	var orderId string

	u, _ := uuid.NewRandom()
	orderId = u.String()
	// 构造交易日志
	payReq := &payment.ChargeReq{
		UserId:  req.UserId,
		OrderId: orderId,
		Amount:  total,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
		},
	}

	// 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{UserId: req.UserId})
	if err != nil {
		klog.Error(err.Error())
	}

	// 支付
	paymentResult, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}

	// 序列化信息
	data, _ := proto.Marshal(&email.EmailReq{
		From:        "from@example.com",
		To:          req.Email,
		ContentType: "text/plain",
		Subject:     "Order subject",
		Content:     "Order Conetne",
	})

	// 构造消息
	msg := &nats.Msg{
		Subject: "email",
		Data:    data,
	}

	// opentelemetry
	// checkout ---{msg.Header}---> opentememetry
	// checkout ---{msg}---> NATS -------> notify
	// notify   ---{msg.Header}---> opentememetry
	// 通过msg.Header形成链路
	otel.GetTextMapPropagator().Inject(s.ctx, propagation.HeaderCarrier(msg.Header))

	// 发送
	_ = mq.Nc.PublishMsg(msg)

	klog.Info(paymentResult)

	resp = &checkout.CheckoutResp{
		OrderId:       orderId,
		TransactionId: paymentResult.TransactionId,
	}
	return
}

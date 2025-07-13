/*
 * @Description: 
 * @version: 
 * @Author: Adxiong
 * @Date: 2025-07-13 17:25:30
 * @LastEditors: Adxiong
 * @LastEditTime: 2025-07-14 00:56:12
 */
package payment

import "adxiong-common-lib/src/internal/service/payment/wechatpay"
import "adxiong-common-lib/src/internal/service/payment/alipay"

struct PaymentContext struct{
	paymentType PayType
	Strategy PaymentStrategy
}

func NewPaymentContext(paymentType PayType) *PaymentContext {
	return &PaymentContext{
		paymentType: paymentType,
		strategy: NewStrategy(paymentType),
	}
}

func (p *PaymentContext) GetPaymentType() PayType {
	return p.paymentType
}

func (p *PaymentContext) CreatePay(req *PayRequest) (*PayResponse, error) {
	return p.strategy.CreatePay(req)
}

func (p *PaymentContext) QueryPay(query *PayQuery) (*PayResponse, error) {
	return p.strategy.QueryPay(query)
}

func (p *PaymentContext) ClosePay(payID string) error {
	return p.strategy.ClosePay(payID)
}

func (p *PaymentContext) Refund(req *RefundRequest) error {
	return p.strategy.Refund(req)
}

func (p *PaymentContext) QueryRefund(refundID string) (*RefundRequest, error) {
	return p.strategy.QueryRefund(refundID)
}

func (p *PaymentContext) VerifyNotification(data []byte) (*PayResponse, error) {
	return p.strategy.VerifyNotification(data)
}


func NewStrategy(payType PayType) (PaymentStrategy,error) {
	switch payType {
	case PayType_WECHAT:
		return &wechatpay.WechatPayStrategy{},nil
	case PayType_ALIPAY:
		return &alipay.AlipayStrategy{},nil
	default:
		return nil,errors.New("unsupported payment type")
	}
}


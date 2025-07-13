/*
 * @Description: 
 * @version: 
 * @Author: Adxiong
 * @Date: 2025-07-13 17:25:21
 * @LastEditors: Adxiong
 * @LastEditTime: 2025-07-14 00:45:07
 */
package payment

// 定义支付接口策略模式
type PaymentStrategy interface {
	// 创建支付
	CreatePay(req *PayRequest) (*PayResponse, error)
	// 查询支付
	QueryPay(query *PayQuery) (*PayResponse, error)
	// 关闭支付
	ClosePay(payID string) error
	// 申请退款
	Refund(req *RefundRequest) error
	// 查询退款
	QueryRefund(refundID string) (*RefundRequest, error)
	// 验证支付通知
	VerifyNotification(data []byte) (*PayResponse, error)
}





/**
 * 策略模式: 不同的算法按照统一的标准封装，客户端根据不同的场景，决策使用何种算法
 * 核心
 * 声明标准：定义interface
 * 封装算法：按照标准interface封装分支代码，得到每一个具体策略
 * 构建算法集：每一个具体策略构成策略池子 -> 这就是沉淀的过程
 */
package main

import (
	"fmt"
	"runtime"
)

const (
	// ConstWechatPay 微信支付
	ConstWechatPay = "wechat_pay"
	// ConstAliPayWap 支付宝支付 网页版
	ConstAliPayWap = "AliPayWapwap"
	// ConstBankPay 银行卡支付
	ConstBankPay = "quickbank"
)

// Context上下文
type Context struct {
	// 用户选择的支付方式
	PayType string `json:"pay_type"`
}

// PaymentInterface接口
type PaymentInterface interface {
	Pay(ctx *Context) error    //支付
	Refund(ctx *Context) error //退款
}

// WechatPay微信支付
type WechatPay struct {
}

func (p *WechatPay) Pay(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "使用微信支付...")
	return
}

func (p *WechatPay) Refund(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "使用微信退款...")
	return
}

// AliPayWeb支付宝支付
type AliPayWeb struct {
}

func (p *AliPayWeb) Pay(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "使用支付宝支付...")
	return
}

func (p *AliPayWeb) Refund(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "使用支付宝退款...")
	return
}

// 银行卡支付
type BankPay struct {
}

func (p *BankPay) Pay(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "使用银行卡支付...")
	return
}

func (p *BankPay) Refund(ctx *Context) (err error) {
	fmt.Println(runFuncName(), "使用银行卡退款...")
	return
}

// 获取正在运行的函数名
func runFuncName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func main() {
	// 业务上下文
	ctx := &Context{PayType: "wechat_pay"}

	// 获取支付方式
	var instance PaymentInterface
	switch ctx.PayType {
	case ConstWechatPay:
		instance = &WechatPay{}
	case ConstAliPayWap:
		instance = &AliPayWeb{}
	case ConstBankPay:
		instance = &BankPay{}
	default:
		panic("无效的支付方式")
	}
	//支付
	instance.Pay(ctx)
}

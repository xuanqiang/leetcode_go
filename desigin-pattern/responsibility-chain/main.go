/**
 * 责任链模式
 * 首先把一系列业务按职责划分成不同的对象，接着把这一系列对象构成一个链，
 * 然后在这一系列对象中传递请求对象，直到被处理为止（调整为直到链路结束为止）结束:异常结束，或链路执行完毕结束
 * 优势：
 * 直观：一眼可观的业务调用过程
 * 无限扩展：可无限扩展的业务逻辑
 * 高度封装：复杂业务代码依然高度封装
 * 极易被修改：复杂业务代码下修改代码只需要专注对应的业务类(结构体)文件即可，以及极易被调整的业务执行顺序
 * 使用场景：
 * 任何杂乱无章的业务代码，都可以使用责任链模式(改)去重构、设计
 */

package main

import (
	"fmt"
	"runtime"
)

// Context context 上下文信息
type Context struct {
}

// Handler处理
type Handler interface {
	// 自身的业务
	Do(c *Context) error
	// 设置下一个对象
	SetNext(h Handler) Handler
	// 执行
	Run(c *Context) error
}

// Next抽象出来的， 可被合成复用的结构体
type Next struct {
	// 下一个对象
	nextHandler Handler
}

//SetNext 可被复用，返回值是下一个对象
func (n *Next) SetNext(h Handler) Handler {
	n.nextHandler = h
	return h
}

// Run执行
func (n *Next) Run(c *Context) (err error) {
	// 由于go无继承的概念 这里无法执行当前handler的Do
	// n.Do(c)
	if n.nextHandler != nil {
		// 执行下一个handler的Do
		if err = (n.nextHandler).Do(c); err != nil {
			return
		}
		// 执行下一个handler的Run
		return (n.nextHandler).Run(c)
	}
	return
}

// NullHandler 空Handler
// 由于go无继承的概念 作为链式调用的第一个载体 设置实际的是下一个对象
type NullHandler struct {
	// 合成复用Next的`nextHandler`成员属性、`SetNext`成员方法、`Run`成员方法
	Next
}

// Do 空handler的Do
func (h *NullHandler) Do(c *Context) (err error) {
	// 空Handler 这里什么也不做 只是载体 do nothing...
	return
}

// ArgumentsHandler 校验参数的handler
type ArgumentsHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *ArgumentsHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "校验参数成功...")
	return
}

// AddressInfoHandler 地址信息handler
type AddressInfoHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *AddressInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "校验参数成功...")
	return
}

// CartInfoHandler 获取购物车数据handler
type CartInfoHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *CartInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取购物车数据...")
	return
}

// StockInfoHandler 商品库存handler 获取商品库存信息
type StockInfoHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *StockInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取商品库存信息...")
	fmt.Println(runFuncName(), "商品库存校验...")
	return
}

// PromotionInfoHandler 获取优惠信息handler
type PromotionInfoHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *PromotionInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取优惠信息...")
	return
}

// ShipmentInfoHandler 获取运费信息handler
type ShipmentInfoHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *ShipmentInfoHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "获取运费信息...")
	return
}

// PromotionUseHandler 使用优惠信息handler
type PromotionUseHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *PromotionUseHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "使用优惠信息...")
	return
}

// StockSubtractHandler 库存操作handler
type StockSubtractHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *StockSubtractHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "扣库存...")
	return
}

// CartDelHandler 清理购物车handler
type CartDelHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *CartDelHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "清理购物车...")
	// err = fmt.Errorf("CartDelHandler.Do fail")
	return
}

// DBTableOrderHandler 写订单表handler
type DBTableOrderHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *DBTableOrderHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "写订单表...")
	return
}

// DBTableOrderSkusHandler 写订单商品表handler
type DBTableOrderSkusHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *DBTableOrderSkusHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "写订单商品表...")
	return
}

// DBTableOrderPromotionsHandler 写订单优惠信息表handler
type DBTableOrderPromotionsHandler struct {
	// 合成复用Next
	Next
}

// Do 校验参数的逻辑
func (h *DBTableOrderPromotionsHandler) Do(c *Context) (err error) {
	fmt.Println(runFuncName(), "写订单优惠信息表...")
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
	// 初始化空handler
	nullHandler := &NullHandler{}
	nullHandler.SetNext(&ArgumentsHandler{}).
		SetNext(&AddressInfoHandler{}).
		SetNext(&CartInfoHandler{}).
		SetNext(&StockInfoHandler{}).
		SetNext(&PromotionInfoHandler{}).
		SetNext(&ShipmentInfoHandler{}).
		SetNext(&PromotionUseHandler{}).
		SetNext(&StockSubtractHandler{}).
		SetNext(&CartDelHandler{}).
		SetNext(&DBTableOrderHandler{}).
		SetNext(&DBTableOrderSkusHandler{}).
		SetNext(&DBTableOrderPromotionsHandler{})
	// 开始执行业务
	if err := nullHandler.Run(&Context{}); err != nil {
		// 异常
		fmt.Println("Fail | Error:" + err.Error())
		return
	}
	// 成功
	fmt.Println("Success")
	return
}

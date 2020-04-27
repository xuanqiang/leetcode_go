package main

// BehaviorOrderCreateInterface 定义一个订单创建行为的接口
type BehaviorOrderCreateInterface interface {
	// 自身的业务
	Do(robot *RobotOrderCreate) error
}

// 1、Param 参数校验
type Param struct {
}

// Do 参数校验
func (behavior *Param) Do(robot *RobotOrderCreate) error {
	// code
	return nil
}

// 2、Address地址校验
type Address struct {
}

// Do 地址校验
func (behavior *Address) Do(robot *RobotOrderCreate) error {
	//code
	return nil
}

// 3、Check其他校验
type Check struct {
}

// Do 其他校验
func (behavior *Check) Do(robot *RobotOrderCreate) error {
	//code
	return nil
}

// 4、Order写订单表
type Order struct {
}

// Do 写订单表
func (behavior *Order) Do(robot *RobotOrderCreate) error {
	//code
	return nil
}

// 5、OrderItem写订单商品信息
type OrderItem struct {
}

// Do 写订单商品信息
func (behavior *OrderItem) Do(robot *RobotOrderCreate) error {
	//code
	return nil
}

// 6、 Log日志
type Log struct {
}

func (behavior *Log) Do(robot *RobotOrderCreate) error {
	//code
	return nil
}

// Request 请求数据对象
type Request struct {
}

// RobotOrderCreate 一个具体的创建订单的机器人
type RobotOrderCreate struct {
	InfoUser     interface{}
	InfoProduct  interface{}
	InfoAddress  interface{}
	InfoCoupon   interface{}
	behaviorList []BehaviorOrderCreateInterface
}

// Init 根据请求参数、购物车数据 初始化机器人的属性
func (robot *RobotOrderCreate) Init(r *Request) *RobotOrderCreate {
	// code
	return robot
}

// RegisterBehavior 注册行为
// @param 行为列表
func (robot *RobotOrderCreate) RegisterBehavior(behavior ...BehaviorOrderCreateInterface) *RobotOrderCreate {
	robot.behaviorList = append(robot.behaviorList, behavior...)
	return robot
}

// Create创建订单
func (robot *RobotOrderCreate) Create() error {
	//code
	for _, behavior := range robot.behaviorList {
		behavior.Do(robot)
	}
	return nil
}

func main() {
	(&RobotOrderCreate{}).Init(&Request{}).
		RegisterBehavior(
			&Param{},
			&Address{},
			&Check{},
			&Order{},
			&OrderItem{},
			&Address{},
			&Log{},
		).
		Create()
}

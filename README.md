# Installation
```
go get github.com/gfandada/gserver
```   
# 规范
```
由于gserver保证性能和健壮性，所以开发者可以把精力放到如何处理消息上，以下是gserver的handler规范
type MsgHandler func ([]interface{})[]interface{}
参数介绍:arg[0].(*network.RawMessage)
	    arg[1].(*service.Session)

type Test struct {
	a int
}

// 1.正常回复客户端
func GetHandler(arg []interface{}) []interface{} {
	// ..........
	return []interface{}{network.RawMessage{
		MsgId: uint16(2001),
		MsgData: &GetAck{
			Test: proto.String("hello service-game"),
		},
	}}
}

// 2.不需要回复客户端
func GetHandler(arg []interface{}) []interface{} {
	// ..........
	return nil
}

// 3.显式终止：自定义错误码（9527表示金币不足）
func GetHandler(arg []interface{}) []interface{} {
	// ..........
	if xxxx {
		panic(9527)
		// 客户端会自动收到错误：errid:9527 errstr:""
	}
	return nil
}

// 4.隐式终止：
func GetHandler(arg []interface{}) []interface{} {
	// ..........
	test := new(Test)
	test = nil
	test.a = 23
	// 运行时错误，空指针
	// 客户端会自动收到错误:
	// errid:1 errstr:"runtime error: invalid memory address or nil pointer dereference"
	return nil
}

// 5.同步更新session：
// 因为用户数据被设计成一个map[string]interface{}可以满足任何形式的用户数据写入
func GetHandler(arg []interface{}) []interface{} {
	// ....
	ret := network.RawMessage{
		MsgId: uint16(2001),
		MsgData: &GetAck{
			Test: proto.String("hello service-game"),
		},
	}
	// 获取or修改sess.....
	return []interface{}{ret}
}
```
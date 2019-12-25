package msg

import (
	// "github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
)

// var Processor network.Processor

var Processor = json.NewProcessor()

func init() {
    // 这里我们注册了一个 JSON 消息 Hello
    Processor.Register(&Hello{})
    Processor.Register(&UserLogin{})
}
// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
type Hello struct {
    Name string
    PWD string
}
// 用户登陆协议
type UserLogin struct {
    LoginName string // 用户名
    LoginPW   string // 密码
}
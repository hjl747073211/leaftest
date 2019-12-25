package internal

import (
    "github.com/name5566/leaf/log"
    "github.com/name5566/leaf/gate"
    "reflect"
    "server/msg"
)
 
func init() {
    // 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
    handler(&msg.Hello{}, handleHello)
}
 
func handler(m interface{}, h interface{}) {
    skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}
 
func handleHello(args []interface{}) {
    // 收到的 Hello 消息
    m := args[0].(*msg.Hello)
    // 消息的发送者
    a := args[1].(gate.Agent)
 
    // 输出收到的消息的内容
    log.Debug("hello %v", m.Name)
    log.Debug("hello %v", m.PWD)
   
 
    // 给发送者回应一个 Hello 消息
    a.WriteMsg(&msg.Hello{
        Name: "client",
    })


        // 收到的 Test 消息
    // m := args[0].(*msg.UserLogin)
    // // 消息的发送者
    // a := args[1].(gate.Agent)
    // // 1 查询数据库--判断用户是不是合法
    // // 2 如果数据库返回查询正确--保存到缓存或者内存
    // // 输出收到的消息的内容
    // log.Debug("Test login %v", m.LoginName)
    // // 给发送者回应一个 Test 消息
    // a.WriteMsg(&msg.UserLogin{
    //     LoginName: "client",
    //     LoginPW: "client",
    // })
}
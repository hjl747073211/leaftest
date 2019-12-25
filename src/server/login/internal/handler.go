package internal

import (
	"bytes"
	"encoding/gob"
	"fmt"

	// "net"
	// "encoding/json"
	// "fmt"

	"reflect"
	"server/msg"

	"github.com/name5566/leaf/db/redis"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	// 向当前模块（login 模块）注册 Protocol.UserLogin 消息的消息处理函数 handleTest
	handleMsg(&msg.UserLogin{}, handleTest)
}

// 消息处理
func handleTest(args []interface{}) {
	// 收到的 Test 消息
	m := args[0].(*msg.UserLogin)
	// 消息的发送者
	a := args[1].(gate.Agent)
	// 1 查询数据库--判断用户是不是合法
	// 2 如果数据库返回查询正确--保存到缓存或者内存
	// 输出收到的消息的内容

	log.Debug("name %v", m.LoginName)
	log.Debug("用户ip %v", args[1].(gate.Agent).RemoteAddr())

	//redis取出数据
	conn, _ := redis.Dial("tcp", ":6379")
	defer conn.Close()

	// ip_json, _ := conn.Do("get", "ip")
	// var onlineConns map[string]gate.Agent
	// if ip_json != nil {
	// 	err := json.Unmarshal([]byte(ip_json).(gate.Agent), &onlineConns)
	// 	log.Debug("%v", onlineConns)
	// }
	// var onlineConns map[string]gate.Agent
	// addr := fmt.Sprintf("%s", args[1].(gate.Agent).RemoteAddr())
	// log.Debug("%T", addr)

	// onlineConns[addr] = a
	// log.Debug("%T", onlineConns)
	// if b, err := json.Marshal(onlineConns); err == nil {
	// 	conn.Do("set", "ip", string(b))
	// 	log.Debug("%T", string(b))
	// }

	// //ips转呈数组
	// log.Debug("ips %v", &onlineConns)
	// if onlineConns != nil {
	// 	addr := fmt.Sprintf("%s", args[1].(gate.Agent).RemoteAddr())
	// 	// if conn, ok := onlineConns[addr]; ok {
	// 	// 	log.Debug("cunzai  %v", conn)
	// 	// } else {
	// 	// 	log.Debug("bucunzai  %v", args[1].(gate.Agent))
	// 	// 	onlineConns[addr] = args[1].(gate.Agent)
	// 	// }
	// 	onlineConns[addr] = args[1].(gate.Agent)
	// }

	// conn.Do("set", "ip", onlineConns)
	// log.Debug("onlineconn %v", onlineConns)
	// 给发送者回应一个 Test 消息
	// onlineConns = make(map[string]gate.Agent)
	// class_map := make(map[string]gate.Agent)

	// xx, _ := redis.String(conn.Do("get", "ip"))
	// log.Debug("%T", xx)
	// log.Debug("%s", xx)
	// _ = json.Unmarshal([]byte(xx), &class_map)

	// log.Debug("%T", class_map)
	// if err != nil {
	// 	log.Debug("str to map err %s", err)
	// }

	// addr := fmt.Sprintf("%s", args[1].(gate.Agent).RemoteAddr())
	// if xx != nil {
	// 	xx[addr] = a
	// 	conn.Do("set", "ip", xx)
	// } else {
	// 	onlineConns = make(map[string]gate.Agent)
	// 	onlineConns[addr] = a
	// 	conn.Do("set", "ip", onlineConns)
	// }

	// onlineConns = make(map[string]interface{})
	aa = make(map[string]interface{})

	// onlineConns[addr] = a
	// log.Debug("%T", onlineConns)
	// conn.Do("set", "ip", onlineConns)

	// addr := fmt.Sprintf("%s", args[1].(gate.Agent).RemoteAddr())
	// onlineConns[addr] = a
	// log.Debug("%s", a)

	// conn.Do("set", "ip", onlineConns)

	// log.Debug("onlineconn %s", onlineConns)
	// xx, _ := redis.Bytes(conn.Do("get", "ip"))
	// log.Debug("%T", xx)
	// log.Debug("%s", xx)
	// dec := gob.NewDecoder(bytes.NewReader(xx))
	// dec.Decode(&onlineConns)
	// log.Debug("onlineconn1 %s", onlineConns[addr])

	// onlineConns[addr].(gate.Agent).WriteMsg(&msg.UserLogin{
	// 	LoginName: m.LoginName,
	// 	LoginPW:   m.LoginPW,
	// })

	//--------------------------------------
	addr := fmt.Sprintf("%s", args[1].(gate.Agent).RemoteAddr())
	xx, _ := redis.Bytes(conn.Do("get", "ip"))
	// if xx != nil {
	dec := gob.NewDecoder(bytes.NewReader(xx))
	dec.Decode(&aa)
	log.Debug("onlineconn1 %v", aa)
	aa[addr] = a
	conn.Do("set", "ip", aa)
	// } else {
	// 	onlineConns = make(map[string]interface{})
	// 	onlineConns[addr] = a
	// 	conn.Do("set", "ip", onlineConns)
	// }

}

var onlineConns map[string]interface{}
var aa map[string]interface{}

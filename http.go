package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/rpc"
)

//等待节点访问
func (rf *Raft) getRequest(writer http.ResponseWriter, request *http.Request) {
	//解析请求参数（表单参数）
	request.ParseForm()
	//http://localhost:8080/req?message=ohmygod
	if len(request.Form["message"]) > 0 && rf.currentLeader != "-1" {
		message := request.Form["message"][0]
		m := new(Message)
		m.MsgID = getRandom()
		m.Msg = message
		//接收到消息后，直接转发到领导者
		fmt.Println("http监听到了消息，准备发送给领导者，消息id:", m.MsgID)
		//领导者的端口号
		port := nodeTable[rf.currentLeader]
		//这里实际上相当于一个客户端
		rp, err := rpc.DialHTTP("tcp", "127.0.0.1"+port)
		if err != nil {
			log.Panic(err)
		}
		//接收服务端返回消息的变量
		b := false
		//调用领导者的消息接收方法
		err = rp.Call("Raft.LeaderReceiveMessage", m, &b)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("消息是否已发送到领导者：", b)
		writer.Write([]byte("ok!!!"))
	}
}

func (rf *Raft) httpListen() {
	//创建getRequest()回调方法
	http.HandleFunc("/req", rf.getRequest)
	fmt.Println("监听8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
		return
	}
}

//返回一个十位数的随机数，作为消息id
func getRandom() int {
	x := big.NewInt(10000000000)
	for {
		result, err := rand.Int(rand.Reader, x)
		if err != nil {
			log.Panic(err)
		}
		if result.Int64() > 1000000000 {
			return int(result.Int64())
		}
	}
}

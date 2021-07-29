package utils

import (
	"chat/commont/message"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
)
func GetTime()string{
	now  := time.Now()
	dateString := fmt.Sprintf("%d-%d-%d %d:%d:%d",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
	return dateString
}

//反序列化聊天内容
func UnMarshalChatData(data string)(string){
	var chatData message.ChatData
	err := json.Unmarshal([]byte(data),&chatData)
	if err != nil{
		fmt.Println("聊天内容反序列化成结构体失败 err:" ,err)
		return "反序列化失败"
	}
	str := fmt.Sprintf("用户【%d】，通过【%s】向你发送了消息：【%s】",chatData.FromId, getchatType(chatData.ChatType) ,chatData.Content)
	return str
}

//聊天方式
func getchatType(chatType int) string {
	if chatType == 2{
		return  "群发"
	}
	return  "私聊"
}


func GetMessage(conn *net.Conn )(m *message.ResMessage){
	buf := make([]byte , 4096)
	n , err  := (*conn).Read(buf)

	if n == 0{ //读取字段尾0、客户端应该下线
		//(*conn).Close()
		return;
	}

	if (err != nil && err != io.EOF){ //读取失败
		fmt.Println("服务器读取客户端的消息失败 ！ err = " ,err)
		return
	}

	//fmt.Println("接受到的消息是：" ,string(buf[:n]))
	var msg message.ResMessage

	//提取消息
	err = json.Unmarshal(buf[:n],&msg)
	if err != nil{
		fmt.Println("接收数据反序列化成结构体失败 err:" ,err)
		return;
	}
	m = &msg
	return
}


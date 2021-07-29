package message

import (
	"encoding/json"
	"fmt"
)

//客户端请求的消息
type ReqMessage struct {
	Type int  `json:"messtype"` //1.登陆  2.交互信息  3.退出
	MsgData string  `json:"msgData"` //消息 json
}

//服务器返回的消息
type ResMessage struct{
	Type int  `json:"messtype"` //1.登陆  2.交互信息  3.退出
	MsgData string  `json:"msgData"` //消息 json
	Code int `json:"code"` //返回的结果  ---    1：请求成功   -1：失败   2:别人发送消息(主推)
}

// 聊天消息结构体
type ChatData struct {
	ChatType int  `json:"chatType"` //1:私聊   2:群聊
	ToId int `json:"toid"`//仅私聊时候对应的id
	Content string `json:"content"`//具体信息
	FromId int `json:"fromId"` //发送聊天的人
}




//创建请求消息体
func CreateReqdMessage(messType int ,msgData string )*ReqMessage{
	sendMes := ReqMessage{
		Type: messType,
		MsgData: msgData,
	}
	return &sendMes
}

//创建回复消息体
func CreateResMessage(messType int ,msgData string ,code int)*ResMessage{
	sendMes := ResMessage{
		Type: messType,
		MsgData: msgData,
		Code: code,
	}
	return &sendMes
}



//构建聊天内容实体
func CreateNewMsg(chatype ,toId int ,content string)*ChatData{
	msg := ChatData{
		ChatType: chatype,
		ToId: toId,
		//FromId: fromId,
		Content: content,
	}
	return &msg;
}


//封包
func (this *ReqMessage) Packet()( []byte , error ){
	reqMess ,err := json.Marshal(this)
	if err != nil{
		fmt.Println("封包失败 ，err: " ,err)
		return nil,err
	}
	return json.Marshal(reqMess)
}



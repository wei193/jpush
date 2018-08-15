package jpush

import (
	"encoding/json"
)

//Message Message
type Message struct {
	MsgContent  string            `json:"msg_content"`
	Title       string            `json:"title,omitempty"`
	ContentType string            `json:"content_type,omitempty"`
	Extras      map[string]string `json:"extras,omitempty"`
}

//SmsMessage SmsMessage
type SmsMessage struct {
	DelayTime int         `json:"delay_time"`
	TempID    int64       `json:"temp_id"`
	TempPara  interface{} `json:"temp_para,omitempty"`
}

//CreateMessage 创建消息
func CreateMessage() (message *Message) {
	return &Message{}
}

//SetMessageContent 设置消息内容
func (message *Message) SetMessageContent(content interface{}) {
	buf, _ := json.Marshal(content)
	message.MsgContent = string(buf)
}

//AddExtras 添加扩展字段
func (message *Message) AddExtras(key, val string) {
	if message.Extras == nil {
		message.Extras = make(map[string]string)
	}
	message.Extras[key] = val
}

package jpush

//Message Message
type Message struct {
	MsgContent  string `json:"msg_content"`
	Title       string `json:"title,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Extras      string `json:"extras,omitempty"`
}

//SmsMessage SmsMessage
type SmsMessage struct {
	DelayTime int         `json:"delay_time"`
	TempID    int64       `json:"temp_id"`
	TempPara  interface{} `json:"temp_para,omitempty"`
}

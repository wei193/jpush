package jpush

import (
	"errors"
	"log"
)

//平台
const (
	Android  = "android"
	IOS      = "ios"
	WinPhone = "winphone"
)

//Push 推送对象结构
type Push struct {
	Platform     interface{} `json:"platform"`
	Audience     interface{} `json:"audience"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	SmsMessage   interface{} `json:"sms_message,omitempty"`
	Options      interface{} `json:"options,omitempty"`
	Cid          interface{} `json:"cid,omitempty"`
}

//Push 推送
func (c *JPush) Push(push Push) (err error) {
	if push.Platform == nil {
		push.Platform = "all"
	}
	if push.Audience == nil {
		push.Audience = "all"
	}

	req, err := createRequset(c.ServerAddr+PushURL, "POST", nil, push)
	if err != nil {
		return err
	}
	buf, err := c.requset(req)
	if err != nil {
		return err
	}
	log.Println(string(buf))
	return
}

// PushNotification 推送通知消息
func (c *JPush) PushNotification(platform *Platform, audience *Audience, notification *Notification, options *Options) (err error) {
	if notification == nil {
		return errors.New("通知消息不能为空")
	}
	push := Push{
		Platform:     platform,
		Audience:     audience,
		Options:      options,
		Notification: notification,
	}
	return c.Push(push)
}

//PushNotificationToTags 推送通知消息到指定标签
func (c *JPush) PushNotificationToTags(tags []string, notification *Notification) (err error) {
	audience := CreateAudience()
	audience.SetTag(tags)
	return c.PushNotification(nil, audience, notification, nil)
}

//PushNotificationToAlias 推送通知消息到指定别名
func (c *JPush) PushNotificationToAlias(alias []string, notification *Notification) (err error) {
	audience := CreateAudience()
	audience.SetAlias(alias)
	return c.PushNotification(nil, audience, notification, nil)
}

//PushMessage 推送自定义消息
func (c *JPush) PushMessage(platform *Platform, audience *Audience, message *Message, options *Options) (err error) {
	if message == nil {
		return errors.New("自定义消息不能为空")
	}
	push := Push{
		Platform: platform,
		Audience: audience,
		Options:  options,
		Message:  message,
	}
	return c.Push(push)
}

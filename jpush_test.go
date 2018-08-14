package jpush

import (
	"log"
	"testing"
	"time"
)

func TestJpush(t *testing.T) {
	p := NewJPush("https://bjapi.push.jiguang.cn", "d4ea42907c4b10c3abd2e440", "baba4a4009bce1f87f03936c")
	notification := CreateNotification("这是一个测试")
	notification.AddExtras("type", "1")
	notification.AddExtras("content", "12")
	err := p.Push(Push{
		Notification: notification,
	})
	if err != nil {
		log.Println(err)
	}
	time.Sleep(time.Second)
}

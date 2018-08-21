package jpush

import (
	"errors"
	"math/rand"
	"time"
)

//Options Options
type Options struct {
	Sendno          int    `json:"sendno,omitempty"`
	TimeToLive      int    `json:"time_to_live,omitempty"`
	OverrideMsgID   int64  `json:"override_msg_id,omitempty"`
	ApnsProduction  bool   `json:"apns_production"`
	ApnsCollapseID  string `json:"apns_collapse_id,omitempty"`
	BigPushDuration int    `json:"big_push_duration,omitempty"`
}

//CreateOptions 创建Options
func CreateOptions() (options *Options) {
	return &Options{ApnsProduction: true}
}

//SentRandSendNo 设置随机INT
func (options *Options) SentRandSendNo() {
	rand.Seed(time.Now().UnixNano())
	options.Sendno = rand.Int()
}

//SetTimeLive 设置离线消息保留时长
func (options *Options) SetTimeLive(time time.Time) (err error) {
	if time.Unix() > 864000 {
		return errors.New("时长错误")
	}
	options.TimeToLive = int(time.Unix())
	return nil
}

//SetApnsProduction APNs是否生产环境
func (options *Options) SetApnsProduction(production bool) {
	options.ApnsProduction = production
}

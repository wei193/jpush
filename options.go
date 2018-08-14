package jpush

//Options Options
type Options struct {
	Sendno          int    `json:"sendno,omitempty"`
	TimeToLive      int    `json:"time_to_live,omitempty"`
	OverrideMsgID   int64  `json:"override_msg_id,omitempty"`
	ApnsProduction  bool   `json:"apns_production,omitempty"`
	ApnsCollapseID  string `json:"apns_collapse_id,omitempty"`
	BigPushDuration int    `json:"big_push_duration,omitempty"`
}

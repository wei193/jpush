package jpush

//Notification 通知
type Notification struct {
	Alert   string               `json:"alert"`
	Android *AndroidNotification `json:"android,omitempty"`
	IOS     *IOSNotification     `json:"ios,omitempty"`
}

//AndroidNotification 安卓推送
type AndroidNotification struct {
	Alert      string            `json:"alert"`
	Title      string            `json:"title,omitempty"`
	BuilderID  int               `json:"builder_id,omitempty"`
	Priority   int               `json:"priority,omitempty"`
	Category   string            `json:"category,omitempty"`
	Style      int               `json:"style,omitempty"`
	AlertType  int               `json:"alert_type,omitempty"`
	BigText    string            `json:"big_text,omitempty"`
	Inbox      interface{}       `json:"inbox,omitempty"`
	BigPicPath string            `json:"big_pic_path,omitempty"`
	Extras     map[string]string `json:"extras,omitempty"`
}

//IOSNotification IOSNotification
type IOSNotification struct {
	Alert            interface{}       `json:"alert"`
	Sound            string            `json:"sound,omitempty"`
	Badge            int               `json:"badge,omitempty"`
	ContentAvailable bool              `json:"content-available,omitempty"`
	MutableContent   bool              `json:"mutable-content,omitempty"`
	Category         string            `json:"category,omitempty"`
	Extras           map[string]string `json:"extras,omitempty"`
}

// IOSAlert IOSAlert
type IOSAlert struct {
	Title        string   `json:"alert"`
	Body         string   `json:"body"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
}

//CreateNotification 创建一个推送
func CreateNotification(alert string) (notification *Notification) {
	return &Notification{
		Alert: alert,
	}
}

//AddAndroidNotification 添加安卓通知
func (notification *Notification) AddAndroidNotification(Android *AndroidNotification) {
	notification.Android = Android
}

//SetAndroidNotificationTitle 设置安卓通知标题
func (notification *Notification) SetAndroidNotificationTitle(title string) {
	if notification.Android == nil {
		notification.Android = &AndroidNotification{
			Alert: notification.Alert,
		}
	}
	notification.Android.Title = title
}

//SetAndroidNotificationBuilder 设置安卓通知栏样式ID
func (notification *Notification) SetAndroidNotificationBuilder(id int) {
	if notification.Android == nil {
		notification.Android = &AndroidNotification{
			Alert: notification.Alert,
		}
	}
	notification.Android.BuilderID = id
}

//SetAndroidNotificationPriority 设置安卓通知栏展示优先级
func (notification *Notification) SetAndroidNotificationPriority(priority int) {
	if notification.Android == nil {
		notification.Android = &AndroidNotification{
			Alert: notification.Alert,
		}
	}
	notification.Android.Priority = priority
}

//SetAndroidNotificationExtras 设置安卓通知扩展字段
func (notification *Notification) SetAndroidNotificationExtras(extras map[string]string) {
	if notification.Android == nil {
		notification.Android = &AndroidNotification{
			Alert: notification.Alert,
		}
	}
	notification.Android.Extras = extras
}

//AddAndroidNotificationExtras 设置安卓通知扩展字段
func (notification *Notification) AddAndroidNotificationExtras(key, val string) {
	if notification.Android == nil {
		notification.Android = &AndroidNotification{
			Alert: notification.Alert,
		}
	}
	if notification.Android.Extras == nil {
		notification.Android.Extras = make(map[string]string)
	}
	notification.Android.Extras[key] = val
}

//AddIOSNotification 添加IOS通知
func (notification *Notification) AddIOSNotification(IOS *IOSNotification) {
	notification.IOS = IOS
}

//SetIOSNotificationExtras 设置IOS通知扩展字段
func (notification *Notification) SetIOSNotificationExtras(extras map[string]string) {
	if notification.IOS == nil {
		notification.IOS = &IOSNotification{
			Alert: notification.Alert,
		}
	}
	notification.IOS.Extras = extras
}

//AddIOSNotificationExtras 设置IOS通知扩展字段
func (notification *Notification) AddIOSNotificationExtras(key, val string) {
	if notification.IOS == nil {
		notification.IOS = &IOSNotification{
			Alert: notification.Alert,
		}
	}
	if notification.IOS.Extras == nil {
		notification.IOS.Extras = make(map[string]string)
	}
	notification.IOS.Extras[key] = val
}

//AddExtras 添加通知扩展
func (notification *Notification) AddExtras(key, val string) {
	if notification.Android == nil {
		notification.Android = &AndroidNotification{
			Alert: notification.Alert,
		}
	}
	if notification.IOS == nil {
		notification.IOS = &IOSNotification{
			Alert: notification.Alert,
		}
	}
	if notification.Android.Extras == nil {
		notification.Android.Extras = make(map[string]string)
	}
	notification.Android.Extras[key] = val

	if notification.IOS.Extras == nil {
		notification.IOS.Extras = make(map[string]string)
	}
	notification.IOS.Extras[key] = val
}

//SetExtras 设置通知扩展
func (notification *Notification) SetExtras(extras map[string]string) {
	if notification.Android == nil {
		notification.Android = &AndroidNotification{
			Alert: notification.Alert,
		}
	}
	if notification.IOS == nil {
		notification.IOS = &IOSNotification{
			Alert: notification.Alert,
		}
	}

	notification.Android.Extras = extras
	notification.IOS.Extras = extras
}

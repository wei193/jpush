package jpush

//Platform 推送平台
type Platform struct {
	OS []string
}

//CreatePlatform 创建平台
func CreatePlatform() (platform *Platform) {
	return &Platform{}
}

//AddPlatform 增加平台
func (platform *Platform) AddPlatform(os []string) {
	if platform.OS == nil {
		platform.OS = os
	} else {
		platform.OS = append(platform.OS, os...)
	}
}

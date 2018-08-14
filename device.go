package jpush

//Device 设备信息
type Device struct {
	Tags   []string `json:"tags"`
	Alias  string   `json:"alias"`
	Mobile string   `json:"mobile"`
}

//GetDevices 查询设备的别名与标签
func (j *JPush) GetDevices(registrationid string) {

}

//PostDevices 设置设备的别名与标签
func (j *JPush) PostDevices() {

}

//AddDevicesTags 添加设备tags
func (j *JPush) AddDevicesTags() {

}

//RemoveDevicesTags 删除设备tags
func (j *JPush) RemoveDevicesTags() {

}

//SetDevicesAlias 设置设备
func (j *JPush) SetDevicesAlias(registrationid, alias string) {

}

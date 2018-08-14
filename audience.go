package jpush

//Audience 推送设备
type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationID []string `json:"registration_id,omitempty"`
	Segment        []string `json:"segment,omitempty"`
	Abtest         []string `json:"abtest,omitempty"`
}

//CreateAudience 创建推送设备设置
func CreateAudience() *Audience {
	return &Audience{}
}

//AddTag 添加标签OR
func (audience *Audience) AddTag(tags []string) {
	if audience.Tag == nil {
		audience.Tag = tags
	} else {
		audience.Tag = append(audience.Tag, tags...)
	}
}

//SetTag 设置标签OR
func (audience *Audience) SetTag(tags []string) {
	audience.Tag = tags
}

//AddTagAnd 添加标签AND
func (audience *Audience) AddTagAnd(tags []string) {
	if audience.TagAnd == nil {
		audience.TagAnd = tags
	} else {
		audience.TagAnd = append(audience.TagAnd, tags...)
	}
}

//SetTagAnd 设置标签AND
func (audience *Audience) SetTagAnd(tags []string) {
	audience.TagAnd = tags
}

//AddTagNot 添加标签Not
func (audience *Audience) AddTagNot(tags []string) {
	if audience.TagNot == nil {
		audience.TagNot = tags
	} else {
		audience.TagNot = append(audience.TagNot, tags...)
	}
}

//SetTagNot 设置标签Not
func (audience *Audience) SetTagNot(tags []string) {
	audience.TagNot = tags
}

//AddAlias 添加别名
func (audience *Audience) AddAlias(alias []string) {
	if audience.TagNot == nil {
		audience.TagNot = alias
	} else {
		audience.TagNot = append(audience.TagNot, alias...)
	}
}

//SetAlias 设置标签Alias
func (audience *Audience) SetAlias(alias []string) {
	audience.Alias = alias
}

//AddRegistrationID 添加注册ID
func (audience *Audience) AddRegistrationID(registrationid []string) {
	if audience.TagNot == nil {
		audience.RegistrationID = registrationid
	} else {
		audience.RegistrationID = append(audience.RegistrationID, registrationid...)
	}
}

//SetRegistrationID 设置标签Alias
func (audience *Audience) SetRegistrationID(registrationid []string) {
	audience.RegistrationID = registrationid
}

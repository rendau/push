package entities

type Session struct {
	Id    int64 `json:"id"`
	Error error `json:"error"`
}

type SendReqSt struct {
	UsrIds     []int64           `json:"usr_ids"`
	Title      string            `json:"title"`
	Body       string            `json:"body"`
	Data       map[string]string `json:"data"`
	Badge      int               `json:"badge"`
	AndroidTag string            `json:"android_tag"`
}

type CommonMessageSt struct {
	Tokens     []string          `json:"registration_ids"`
	TimeToLive uint64            `json:"time_to_live"`
	Priority   string            `json:"priority,omitempty"`
	Data       map[string]string `json:"data,omitempty"`
}

type CommonNotificationSt struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

type AndroidMessageSt struct {
	CommonMessageSt
	Notification *AndroidNotificationSt `json:"notification,omitempty"`
}

type AndroidNotificationSt struct {
	CommonNotificationSt
	Tag         string `json:"tag,omitempty"`
	Sound       string `json:"sound,omitempty"`
	ClickAction string `json:"click_action,omitempty"`
}

type IosMessageSt struct {
	CommonMessageSt
	Notification *IosNotificationSt `json:"notification,omitempty"`
}

type IosNotificationSt struct {
	CommonNotificationSt
	Sound string `json:"sound,omitempty"`
	Badge string `json:"badge,omitempty"`
}

type UndefinedMessageSt struct {
	CommonMessageSt
	Notification *UndefinedNotificationSt `json:"notification,omitempty"`
}

type UndefinedNotificationSt struct {
	CommonNotificationSt
	Sound       string `json:"sound,omitempty"`
	Badge       string `json:"badge,omitempty"`
	ClickAction string `json:"click_action,omitempty"`
}

type WebMessageSt struct {
	CommonMessageSt
}

type FcmReplySt struct {
	Results []FcmReplyResultSt `json:"results"`
}

type FcmReplyResultSt struct {
	Error string `json:"error"`
}

type TokenCreateSt struct {
	Token      string
	Id         int64
	PlatformId int
}

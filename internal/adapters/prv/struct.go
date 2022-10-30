package prv

type SendReqSt struct {
	PlatformId int
	Tokens     []string
	Title      string
	Body       string
	Data       map[string]string
	Badge      int
	AndroidTag string
}

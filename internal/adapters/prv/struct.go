package prv

type SendReqSt struct {
	Tokens     []string
	Title      string
	Body       string
	Data       map[string]string
	Badge      int
	AndroidTag string
}

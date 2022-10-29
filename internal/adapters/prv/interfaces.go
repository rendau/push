package prv

type Prv interface {
	Send(obj *SendReqSt) error
}

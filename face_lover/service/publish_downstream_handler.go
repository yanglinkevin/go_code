package service

type IHander interface {
	GetInfo(id string, sourceType int16) (result string)
	SendToDownStream(result string)
}



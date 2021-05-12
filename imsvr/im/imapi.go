package im


type ImApi interface {
	SendMsg() error
	GetMsg () error
}
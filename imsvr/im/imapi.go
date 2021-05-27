package im

type ImApi interface {
	SendMsg() error
	GetMsg() error
	DelMsg() error
	EditMsg() error
	PinMsg() error
	RecallMsg() error
}

package app

type GlobalService interface {
	GetGlobalAverage() (GlobalStatus, error)
}

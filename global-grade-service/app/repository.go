package app

type GlobalRepo interface {
	GlobalAverage() (GlobalStatus, error)
}

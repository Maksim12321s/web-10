package api

type Usecase interface {
	AddHello(string) error
	DeleteHello(string, bool) error
	GetQuantity() (int, error)
}

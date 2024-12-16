package api

type Usecase interface {
	AddHello(string) error
	DeleteHello(string, bool) error
	AllHellos() ([]string, error)
}

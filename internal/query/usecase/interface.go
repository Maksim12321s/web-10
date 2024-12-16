package usecase

type Provider interface {
	PostHello(msg string) error
	DeleteHello(msg string, all bool) error
	GetHello() ([]string, error)
}

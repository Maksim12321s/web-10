package usecase

type Usecase struct {
	pr Provider
}

func CreateNewUsecase(pr Provider) *Usecase {
	return &Usecase{
		pr: pr,
	}
}

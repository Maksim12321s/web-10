package usecase

func (use *Usecase) AddHello(msg string) error {
	if err := use.pr.PostHello(msg); err != nil {
		return err
	}
	return nil
}

func (use *Usecase) AllHellos() ([]string, error) {
	k, err := use.pr.GetHello()
	if err != nil {
		return nil, err
	}
	return k, nil
}

func (use *Usecase) DeleteHello(msg string, all bool) error {
	if err := use.pr.DeleteHello(msg, all); err != nil {
		return err
	}
	return nil

}

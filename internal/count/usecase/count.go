package usecase

func (use *Usecase) AddHello(msg string) error {
	if err := use.pr.PostHello(msg); err != nil {
		return err
	}
	return nil
}

func (use *Usecase) GetQuantity() (int, error) {
	k, err := use.pr.GetCount()
	if err != nil {
		return -1, err
	}
	return k, nil
}

func (use *Usecase) DeleteHello(msg string, all bool) error {
	if err := use.pr.DeleteHello(msg, all); err != nil {
		return err
	}
	return nil

}

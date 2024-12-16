package provider

func (b *Provider) GetCount() (int, error) {
	var k int
	row := b.conn.QueryRow(`select kolvo from kol`)
	if err := row.Scan(&k); err != nil {
		return -1, err
	}
	return k, nil
}

func (b *Provider) PostHello(msg string) error {
	if _, err := b.conn.Exec(`insert into hello (message) values ($1)`, msg); err != nil {
		return err
	}
	if _, err := b.conn.Exec(`update kol set kolvo = kolvo + 1`); err != nil {
		return err
	}
	return nil
}

func (b *Provider) DeleteHello(msg string, all bool) error {
	if all {
		if _, err := b.conn.Exec(`truncate hello`); err != nil {
			return err
		}
		if _, err := b.conn.Exec(`update kol set kolvo = 0`); err != nil {
			return err
		}

	} else {
		if _, err := b.conn.Exec(`delete from hello where message = $1`, msg); err != nil {
			return err
		}
		if _, err := b.conn.Exec(`update kol set kolvo = kolvo - 1`); err != nil {
			return err
		}
	}
	return nil
}

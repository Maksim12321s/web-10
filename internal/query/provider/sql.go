package provider

func (b *Provider) GetHello() ([]string, error) {
	var k []string
	row, err := b.conn.Query(`select message from hello`)
	if err != nil {
		return nil, err
	}
	var t string
	for row.Next() {
		if err := row.Scan(&t); err != nil {
			return nil, err
		}
		k = append(k, t)
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

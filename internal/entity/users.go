package entity

import (
	"encoding/json"
	"time"
)

const (
	timeDMY = "02.01.2006"
)

type User struct {
	Nickname string    `json:"nickname"`
	Birthday time.Time `json:"birthday"`
}

func (p *User) UnmarshalJSON(b []byte) error {
	type userAlias User
	alias := &struct {
		*userAlias
		Birthday string `json:"birthday"`
	}{
		userAlias: (*userAlias)(p),
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}
	t, err := parseDate(alias.Birthday)
	if err != nil {
		return err
	}
	p.Birthday = t
	return nil
}

func parseDate(in string) (time.Time, error) {
	return time.Parse(timeDMY, in)
}

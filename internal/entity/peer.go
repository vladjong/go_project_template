package entity

import (
	"encoding/json"
	"time"
)

const (
	timeDMY = "02.01.2006"
)

type Peer struct {
	Nickname string    `json:"nickname"`
	Birthday time.Time `json:"birthday,time.RFC3339"`
}

func (p *Peer) UnmarshalJSON(b []byte) error {
	type peerAlias Peer
	alias := &struct {
		*peerAlias
		Birthday string `json:"birthday"`
	}{
		peerAlias: (*peerAlias)(p),
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

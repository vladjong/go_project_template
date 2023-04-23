package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/entity/dto"
)

const (
	peersTable = "peers"
)

func (r *repository) AddPeer(ctx context.Context, peer dto.Peer) error {
	insertQuery, _, err := goqu.Insert(peersTable).Rows(peer).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}
	fmt.Println(insertQuery)

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetPeers(ctx context.Context) ([]entity.Peer, error) {
	selectQuery, _, err := goqu.From(peersTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	peersDto := []dto.Peer{}
	if err = r.db.DB.SelectContext(ctx, &peersDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	peers := make([]entity.Peer, len(peersDto))
	for i, peerDto := range peersDto {
		peer := entity.Peer{
			Nickname: peerDto.Nickname,
			Birthday: peerDto.Birthday,
		}
		peers[i] = peer
	}
	return peers, nil
}

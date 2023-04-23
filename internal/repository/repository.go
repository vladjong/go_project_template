package repository

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/entity/dto"
)

type Repository interface {
	Peer
}

type Peer interface {
	AddPeer(ctx context.Context, item dto.Peer) error
	GetPeers(ctx context.Context) ([]entity.Peer, error)
}

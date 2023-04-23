package peer

import (
	"context"

	"github.com/vladjong/go_project_template/internal/entity"
)

type Service interface {
	AddPeer(ctx context.Context, item entity.Peer) error
	GetPeers(ctx context.Context) ([]entity.Peer, error)
}

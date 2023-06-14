package peer

import (
	"context"
	"fmt"

	"github.com/vladjong/go_project_template/internal/entity"
	"github.com/vladjong/go_project_template/internal/entity/dto"
	"github.com/vladjong/go_project_template/internal/repository"
)

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) AddPeer(ctx context.Context, peer entity.Peer) error {
	peerDto := dto.Peer{
		Nickname: peer.Nickname,
		Birthday: peer.Birthday,
	}
	if err := s.repo.AddPeer(context.Background(), peerDto); err != nil {
		return fmt.Errorf("add peer: %w", err)
	}
	return nil
}

func (s *service) GetPeers(ctx context.Context) ([]entity.Peer, error) {
	items, err := s.repo.GetPeers(context.Background())
	if err != nil {
		return nil, fmt.Errorf("get peers: %w", err)
	}
	return items, nil
}

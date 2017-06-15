package service

import (
	"github.com/thisiserico/golabox/domain"
	"github.com/thisiserico/golabox/eventbus/publisher"
	"github.com/thisiserico/golabox/readdomain"
)

type Service struct {
	commandRepo    domain.CommandRepository
	queryRepo      domain.QueryRepository
	readClient     readdomain.Repository
	eventPublisher *publisher.Publisher
}

func New(
	cr domain.CommandRepository,
	qr domain.QueryRepository,
	rc readdomain.Repository,
	ep *publisher.Publisher,
) *Service {
	return &Service{
		commandRepo:    cr,
		queryRepo:      qr,
		readClient:     rc,
		eventPublisher: ep,
	}
}

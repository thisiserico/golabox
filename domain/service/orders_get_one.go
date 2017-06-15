package service

import (
	"fmt"

	"github.com/thisiserico/golabox/domain"
)

func (srv *Service) GetOrder(oID domain.OrderId) (*domain.Order, error) {
	order := srv.queryRepo.GetOrder(oID)
	if order == nil {
		return nil, fmt.Errorf("Non existing order %s", oID)
	}

	return order, nil
}

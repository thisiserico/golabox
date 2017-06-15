package service

import (
	"fmt"

	"github.com/thisiserico/golabox/domain"
)

func (srv *Service) GetItem(iID domain.ItemId) (*domain.Item, error) {
	item := srv.queryRepo.GetOrderItem(iID)
	if item == nil {
		return nil, fmt.Errorf("Non existing item %s", iID)
	}

	return item, nil
}

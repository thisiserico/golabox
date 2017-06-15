package service

import "github.com/thisiserico/golabox/domain"

func (srv *Service) CreateNewOrderItem(pID domain.ProductId, quantity int) (*domain.Item, error) {
	item := domain.NewItem(pID, quantity)
	if err := srv.commandRepo.UpsertOrderItem(item); err != nil {
		return nil, err
	}

	return item, nil
}

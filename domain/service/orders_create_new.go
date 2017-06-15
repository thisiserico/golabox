package service

import "github.com/thisiserico/golabox/domain"

func (srv *Service) CreateNewOrder() (*domain.Order, error) {
	order := domain.NewOrder()

	if err := srv.commandRepo.UpsertOrder(order); err != nil {
		return nil, err
	}

	srv.readClient.CreateOrder(order)

	return order, nil
}

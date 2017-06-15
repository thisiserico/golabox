package service

import "github.com/thisiserico/golabox/domain"

func (srv *Service) UpdateItemQuantity(i *domain.Item, quantity int) error {
	i.UpdateQuantity(quantity)

	return srv.commandRepo.UpsertOrderItem(i)
}

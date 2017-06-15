package service

import "github.com/thisiserico/golabox/domain"

func (srv *Service) DeleteItem(iID domain.ItemId) error {
	return srv.commandRepo.DeleteOrderItem(iID)
}

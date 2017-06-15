package service

import "github.com/thisiserico/golabox/domain"

func (srv *Service) AddItemToOrder(o *domain.Order, oi *domain.Item) error {
	o.Items = append(o.Items, oi.ID)

	p := srv.queryRepo.GetProduct(oi.Product)
	srv.readClient.AddItemToOrder(o, oi, p)

	return srv.commandRepo.UpsertOrder(o)
}

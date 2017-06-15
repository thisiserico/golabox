package service

import "github.com/thisiserico/golabox/domain"

func (srv *Service) GetAllProducts() []*domain.Product {
	return srv.queryRepo.GetAllProducts()
}

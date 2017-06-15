package service

import (
	"fmt"

	"github.com/thisiserico/golabox/domain"
)

func (srv *Service) GetProduct(pID domain.ProductId) (*domain.Product, error) {
	if p := srv.queryRepo.GetProduct(pID); p != nil {
		return p, nil
	}

	return nil, fmt.Errorf("Non existing product %s", pID)
}

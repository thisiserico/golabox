package service

import (
	"github.com/thisiserico/golabox/domain"
)

func (srv *Service) DecreaseProductsStock(o *domain.Order) error {
	for _, i := range o.Items {
		item := srv.queryRepo.GetOrderItem(i)
		product := srv.queryRepo.GetProduct(item.Product)

		product.Stock = product.Stock - item.Quantity
		if err := srv.commandRepo.UpsertProduct(product); err != nil {
			return err
		}
	}

	return nil
}

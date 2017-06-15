package service

import (
	"fmt"

	"github.com/thisiserico/golabox/domain"
)

func (srv *Service) PayOrder(o *domain.Order, p *domain.Payment) error {
	if err := srv.validateProductQuantities(o); err != nil {
		return err
	}

	if err := srv.commandRepo.UpsertPayment(p); err != nil {
		return err
	}

	o.Payment = p.ID

	if err := srv.commandRepo.UpsertOrder(o); err != nil {
		return nil
	}

	return srv.DecreaseProductsStock(o)
}

func (srv *Service) validateProductQuantities(o *domain.Order) error {
	for _, i := range o.Items {
		item := srv.queryRepo.GetOrderItem(i)
		product := srv.queryRepo.GetProduct(item.Product)

		if product.Stock < item.Quantity {
			return fmt.Errorf("not enough '%s' (%s) in stock", product.Name, product.ID)
		}
	}

	return nil
}

package controller

import (
	"context"
	"CmsProject/CmsProject/service"
)

type OrderController struct {
	Ctx context.Context
	Service service.OrderService
}

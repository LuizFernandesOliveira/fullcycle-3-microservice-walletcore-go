package gateway

import "github.com/LuizFernandesOliveira/fullcyle-3-microservice-walletcore-go/internal/entity"

type AccountGateway interface {
	Save(client *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}

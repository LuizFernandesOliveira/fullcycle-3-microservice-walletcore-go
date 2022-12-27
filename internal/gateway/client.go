package gateway

import "github.com/LuizFernandesOliveira/fullcyle-3-microservice-walletcore-go/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}

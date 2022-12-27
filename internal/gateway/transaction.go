package gateway

import "github.com/LuizFernandesOliveira/fullcyle-3-microservice-walletcore-go/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}

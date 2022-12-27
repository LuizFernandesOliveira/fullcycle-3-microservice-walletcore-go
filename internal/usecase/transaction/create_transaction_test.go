package transaction

import (
	"github.com/LuizFernandesOliveira/fullcyle-3-microservice-walletcore-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	clientFrom, _ := entity.NewClient("John Doe", "j@j.com")
	accountFrom := entity.NewAccount(clientFrom)
	accountFrom.Credit(100)

	clientTo, _ := entity.NewClient("John Doe To", "l@l.com")
	accountTo := entity.NewAccount(clientTo)
	accountTo.Credit(100)

	accountGateway := &AccountGatewayMock{}
	accountGateway.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGateway.On("FindByID", accountTo.ID).Return(accountTo, nil)

	transactionGateway := &TransactionGatewayMock{}
	transactionGateway.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        10,
	}

	createTransactionUseCase := NewCreateTransactionUseCase(transactionGateway, accountGateway)
	output, err := createTransactionUseCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	accountGateway.AssertExpectations(t)
	accountGateway.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGateway.AssertExpectations(t)
	transactionGateway.AssertNumberOfCalls(t, "Create", 1)
}

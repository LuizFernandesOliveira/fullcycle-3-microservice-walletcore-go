package account

import (
	"github.com/LuizFernandesOliveira/fullcyle-3-microservice-walletcore-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "j@l.com")
	clientGateway := &ClientGatewayMock{}
	clientGateway.On("Get", mock.Anything).Return(client, nil)

	accountGateway := &AccountGatewayMock{}
	accountGateway.On("Save", mock.Anything).Return(nil)

	createAccountUseCase := NewCreateAccountUseCase(accountGateway, clientGateway)
	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := createAccountUseCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientGateway.AssertExpectations(t)
	clientGateway.AssertNumberOfCalls(t, "Get", 1)
	accountGateway.AssertExpectations(t)
	accountGateway.AssertNumberOfCalls(t, "Save", 1)
}

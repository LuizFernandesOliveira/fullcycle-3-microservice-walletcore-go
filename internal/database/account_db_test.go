package database

import (
	"database/sql"
	"github.com/LuizFernandesOliveira/fullcyle-3-microservice-walletcore-go/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (suite *AccountDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance FLOAT, created_at date, updated_at date)")
	suite.accountDB = NewAccountDB(db)
	suite.client, _ = entity.NewClient("John Doe", "j@j.com")
}

func (suite *AccountDBTestSuite) TearDownTest() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (suite *AccountDBTestSuite) TestSave() {
	account := entity.NewAccount(suite.client)
	err := suite.accountDB.Save(account)
	suite.Nil(err)
}

func (suite *AccountDBTestSuite) TestFindById() {
	suite.db.Exec("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)",
		suite.client.ID, suite.client.Name, suite.client.Email, suite.client.CreatedAt, suite.client.UpdatedAt,
	)

	account := entity.NewAccount(suite.client)
	err := suite.accountDB.Save(account)
	suite.Nil(err)
	accountDb, err := suite.accountDB.FindById(account.ID)
	suite.Nil(err)
	suite.NotNil(accountDb)
	suite.Equal(account.ID, accountDb.ID)
	suite.Equal(account.Client.ID, accountDb.Client.ID)
	suite.Equal(account.Balance, accountDb.Balance)
}

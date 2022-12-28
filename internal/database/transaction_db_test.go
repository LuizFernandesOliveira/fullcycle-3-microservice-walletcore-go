package database

import (
	"database/sql"
	"github.com/LuizFernandesOliveira/fullcyle-3-microservice-walletcore-go/internal/entity"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	clientFrom    *entity.Client
	clientTo      *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (suite *TransactionDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at date, updated_at date)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance FLOAT, created_at date, updated_at date)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255), account_id_from VARCHAR(255), account_id_to VARCHAR(255), amount FLOAT, created_at date, updated_at date)")
	suite.clientFrom, _ = entity.NewClient("John Doe", "j@j.com")
	suite.clientTo, _ = entity.NewClient("John Doe", "b@b.com")

	accountFrom := entity.NewAccount(suite.clientFrom)
	accountFrom.Balance = 1000
	suite.accountFrom = accountFrom

	accountTo := entity.NewAccount(suite.clientTo)
	accountTo.Balance = 1000
	suite.accountTo = accountTo

	suite.transactionDB = NewTransactionDB(db)
}

func (suite *TransactionDBTestSuite) TearDownTest() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE transactions")
	suite.db.Exec("DROP TABLE accounts")
	suite.db.Exec("DROP TABLE clients")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (suite *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(suite.accountFrom, suite.accountTo, 100)
	suite.Nil(err)
	err = suite.transactionDB.Create(transaction)
	suite.Nil(err)
}

package database

import (
	"database/sql"
	"github.com/LuizFernandesOliveira/fullcyle-3-microservice-walletcore-go/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (suite *ClientDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at date, updated_at date)")
	suite.clientDB = NewClientDB(db)
}

func (suite *ClientDBTestSuite) TearDownTest() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSave() {
	client := &entity.Client{
		ID:    "123",
		Name:  "John Doe",
		Email: "t@t.com",
	}
	err := suite.clientDB.Save(client)
	suite.Nil(err)
}

func (suite *ClientDBTestSuite) TestGet() {
	client, _ := entity.NewClient("John Doe", "t@t.com")
	suite.clientDB.Save(client)
	clientSaved, err := suite.clientDB.Get(client.ID)
	suite.Nil(err)
	suite.NotNil(clientSaved)
	suite.Equal(client.Name, clientSaved.Name)
	suite.Equal(client.Email, clientSaved.Email)
}

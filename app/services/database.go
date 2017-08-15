package services

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nicosullivan/todo-go/app/models"
	"github.com/revel/revel"
	"os"
)

type Database struct {
	Orm *gorm.DB
}

func ConnectToDatabase() *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("TODOHOST"),
		os.Getenv("TODOUSER"),
		os.Getenv("TODONAME"),
		os.Getenv("TODOSSL"),
		os.Getenv("TODOPASSWORD"))

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)

	return db
}

func InitDB() {
	db := ConnectToDatabase()

	db.AutoMigrate(&models.Todo{})

	db.Close()
}

func (c *Database) Open() revel.Result {
	c.Orm = ConnectToDatabase()
	return nil
}

func (c *Database) Close() revel.Result {
	c.Orm.Close()
	return nil
}

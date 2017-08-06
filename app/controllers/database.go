package controllers

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nicosullivan/todo-go/app/models"
	"github.com/revel/revel"
)

type DatabaseController struct {
	*revel.Controller
	db *gorm.DB
}

func InitDB() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gorm sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	db.AutoMigrate(&models.Todo{})
}

func (c *DatabaseController) Open() revel.Result {
	var err error
	c.db, err = gorm.Open("postgres", "host=localhost user=postgres dbname=gorm sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	c.db.SingularTable(true)
	return nil
}

func (c *DatabaseController) Close() revel.Result {
	c.db.Close()
	return nil
}

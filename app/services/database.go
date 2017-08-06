package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nicosullivan/todo-go/app/models"
	"github.com/revel/revel"
)

type Database struct {
	Orm *gorm.DB
}

func InitDB() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=gorm sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)

	db.AutoMigrate(&models.Todo{})
}

func (c *Database) Open() revel.Result {
	var err error
	c.Orm, err = gorm.Open("postgres", "host=localhost user=postgres dbname=gorm sslmode=disable password=postgres")
	if err != nil {
		panic(err)
	}
	c.Orm.SingularTable(true)
	return nil
}

func (c *Database) Close() revel.Result {
	c.Orm.Close()
	return nil
}

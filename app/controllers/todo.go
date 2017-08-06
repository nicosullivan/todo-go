package controllers

import (
	"github.com/nicosullivan/todo-go/app/models"
	"github.com/nicosullivan/todo-go/app/services"
	"github.com/revel/revel"
)

type Todo struct {
	*revel.Controller
	services.Database
}

func (c Todo) Index() revel.Result {
	todos := make([]models.Todo, 0)

	if err := c.Orm.Find(&todos).Error; err != nil {
		panic(err)
	}
	return c.Render(todos)
}

func (c Todo) Create() revel.Result {
	if err := c.Orm.Create(&models.Todo{Done: false, Title: "Third"}).Error; err != nil {
		panic(err)
	}
	return c.Render()
}

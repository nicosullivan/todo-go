package controllers

import (
	"github.com/nicosullivan/todo-go/app/models"
	"github.com/revel/revel"
)

type Todo struct {
	DatabaseController
}

func (c Todo) Index() revel.Result {
	todos := make([]models.Todo, 0)

	if err := c.db.Find(&todos).Error; err != nil {
		panic(err)
	}
	return c.Render(todos)
}

func (c Todo) Create() revel.Result {
	if err := c.db.Create(&models.Todo{Done: false, Title: "Third"}).Error; err != nil {
		panic(err)
	}
	return c.Render()
}

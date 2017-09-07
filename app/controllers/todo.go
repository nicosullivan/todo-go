package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/nicosullivan/todo-go/app/models"
	"github.com/nicosullivan/todo-go/app/services"
	"github.com/revel/revel"
	"log"
	"net/http"
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

func (c Todo) New() revel.Result {
	var newTodo models.Todo
	response := JsonResponse{Success: true}
	c.Response.Status = http.StatusCreated

	if err := json.Unmarshal(c.Params.JSON, &newTodo); err != nil {
		//TODO: check if no content
		//TODO: check if dates can be set via unmarshal
		c.Response.Status = http.StatusInternalServerError
		response.Success = false
		log.Print(err)
	}

	if c.Orm.NewRecord(newTodo) {
		c.Orm.Create(&newTodo)
		c.Response.Out.Header().Add("Location", fmt.Sprintf("todo/%d", newTodo.ID))
		if err := c.Orm.Error; err != nil {
			c.Response.Status = http.StatusInternalServerError
			response.Success = false
			log.Print(err)
		}
	} else {
		c.Response.Status = http.StatusConflict
		response.Success = false
	}

	return c.RenderJSON(response)
}

func (c Todo) Save(id uint) revel.Result {
	var todo models.Todo
	response := JsonResponse{Success: true}

	todo.ID = id
	c.Orm.Find(&todo)

	if err := json.Unmarshal(c.Params.JSON, &todo); err != nil {
		response.Success = false
		log.Print(err)
	}

	c.Orm.Save(&todo)

	return c.RenderJSON(response)
}

func (c Todo) Delete(id uint) revel.Result {
	var todo models.Todo
	todo.ID = id
	response := JsonResponse{Success: true}

	c.Orm.Delete(&todo)

	if err := c.Orm.Error; err != nil {
		response.Success = false
		log.Print(err)
	}

	return c.RenderJSON(response)
}

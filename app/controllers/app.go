package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	DatabaseController
}

func (c App) Index() revel.Result {
	return c.Render()
}

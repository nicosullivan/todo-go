package controllers

import "github.com/revel/revel"

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*DatabaseController).Open, revel.BEFORE)
	revel.InterceptMethod((*DatabaseController).Close, revel.FINALLY)
}

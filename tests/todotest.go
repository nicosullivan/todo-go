package tests

import (
	"github.com/revel/revel/testing"
)

type TodoTest struct {
	testing.TestSuite
}

func (t *TodoTest) Before() {
	println("Set up")
}

func (t *TodoTest) TestThatIndexPageWorks() {
	t.Get("/todo/index")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *TodoTest) After() {
	println("Tear down")
}

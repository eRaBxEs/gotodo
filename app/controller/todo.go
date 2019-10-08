package controller

import (
	"gotodo/lib/model"
	"net/http"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"go.uber.org/zap"
)

// Todo ...
type Todo struct {
	db  *pg.DB
	env *Environment
	log *zap.SugaredLogger
}

// used by the compiler to confirm if Todo{} conforms to Handler
var _ Handler = &Todo{}

// Init ... Here the Todo is implementing the interface Handler
func (s *Todo) Init(env *Environment, prefix string) error {
	s.env = env
	s.db = env.DB
	s.log = env.Log.Sugar()

	rtr := env.Rtr
	g := rtr.Group(prefix)
	p := g.Group("/todo")

	// for portal
	p.GET("/gettasks", s.GetTodos)
	p.POST("/addtask", s.SaveTodo)
	p.DELETE("/deletetask/:id", s.DeleteTodo)

	return nil
}

// GetTodos ...
func (s *Todo) GetTodos(c echo.Context) error {

	task := model.Task{}

	tasks, err := task.GetAll(s.db)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tasks)
}

// SaveTodo ...
func (s *Todo) SaveTodo(c echo.Context) error {

	task := model.Task{}

	if err := c.Bind(&task); err != nil {
		return err
	}

	if err := task.Save(s.db); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, H{
		"created": task.ID,
	})
}

// DeleteTodo ...
func (s *Todo) DeleteTodo(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	tID := int64(id)

	task := model.Task{ID: tID}

	if err := task.Delete(s.db); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, H{
		"deleted": tID,
	})
}

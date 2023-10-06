package controller

import (
	"strconv"

	"github.com/Zoopast/notes-app/model"
	"github.com/Zoopast/notes-app/service"
	"github.com/gin-gonic/gin"
)

type NotesController interface {
	FindAll(ctx *gin.Context) []model.Note
	Find(ctx *gin.Context) model.Note
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.NoteService
}

func New(service service.NoteService) NotesController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var note model.Note
	err := ctx.ShouldBindJSON(&note)

	if err != nil {
		return err
	}

	c.service.Save(note)

	return nil
}
func (c *controller) Update(ctx *gin.Context) error {
	var note model.Note

	err := ctx.ShouldBindJSON(&note)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}

	note.ID = uint(id)

	c.service.Update(note)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
	var note model.Note
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	note.ID = uint(id)
	c.service.Delete(note)
	return nil
}
func (c *controller) Find(ctx *gin.Context) model.Note {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return model.Note{}
	}
	note := c.service.Find(id)
	return note
}
func (c *controller) FindAll(ctx *gin.Context) []model.Note {

	return c.service.FindAll()
}

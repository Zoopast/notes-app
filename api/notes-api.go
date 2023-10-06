package api

import (
	"net/http"

	"github.com/Zoopast/notes-app/controller"
	"github.com/Zoopast/notes-app/dto"
	"github.com/gin-gonic/gin"
)

type NoteApi struct {
	noteController controller.NoteController
}

func NewNoteAPI(
	noteController controller.NoteController,
) *NoteApi {
	return &NoteApi{
		noteController: noteController,
	}
}

func (api *NoteApi) GetAllNotes(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.noteController.FindAll())
}

func (api *NoteApi) GetNote(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.noteController.Find(ctx))
}

func (api *NoteApi) CreateNote(ctx *gin.Context) {
	err := api.noteController.Save(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			&dto.Response{
				Message: err.Error(),
			})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Note created!",
		})
	}
}

func (api *NoteApi) UpdateNote(ctx *gin.Context) {
	err := api.noteController.Update(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			&dto.Response{
				Message: err.Error(),
			})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Note updated!",
		})
	}
}

func (api *NoteApi) DeleteNote(ctx *gin.Context) {
	err := api.noteController.Delete(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			&dto.Response{
				Message: err.Error(),
			})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Note deleted!",
		})
	}
}

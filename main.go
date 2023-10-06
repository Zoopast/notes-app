package main

import (
	"github.com/Zoopast/notes-app/api"
	"github.com/Zoopast/notes-app/controller"
	"github.com/Zoopast/notes-app/repository"
	"github.com/Zoopast/notes-app/service"
	"github.com/gin-gonic/gin"
)

var (
	noteRepository repository.NoteRepository = repository.NewNoteRepository()
	noteService    service.NoteService       = service.New(noteRepository)

	noteController controller.NoteController = controller.New(noteService)
)

func main() {
	server := gin.Default()
	notesAPI := api.NewNoteAPI(noteController)

	apiRoutes := server.Group("/api/v1")
	{
		notes := apiRoutes.Group("/notes")
		{
			notes.GET("", notesAPI.GetAllNotes)
			notes.GET(":id", notesAPI.GetNote)
			notes.POST("", notesAPI.CreateNote)
			notes.PUT("", notesAPI.UpdateNote)
			notes.DELETE("", notesAPI.DeleteNote)
		}
	}

	server.Run()
}

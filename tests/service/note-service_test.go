package service

import (
	"github.com/Zoopast/notes-app/model"
	"github.com/Zoopast/notes-app/repository"
	"github.com/Zoopast/notes-app/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	TITLE   = "Note title test"
	CONTENT = "Note description test"
)

var testNote model.Note = model.Note{
	Title:   TITLE,
	Content: CONTENT,
}

var _ = Describe("Note service", func() {
	var (
		noteRepository repository.NoteRepository
		noteService    service.NoteService
	)

	BeforeSuite(func() {
		noteRepository = repository.NewNoteRepository("../../")
		noteService = service.New(noteRepository)
	})

	Describe("Fetching all existing notes", func() {
		Context("If there is a note in the database", func() {
			BeforeEach(func() {
				noteService.Save(testNote)
			})

			It("should return at least one element", func() {
				notelist := noteService.FindAll()

				Ω(notelist).ShouldNot(BeEmpty())
			})

			It("should map the fields correctly", func() {
				firstNote := noteService.FindAll()[0]

				Ω(firstNote.Title).Should(Equal(TITLE))
				Ω(firstNote.Content).Should(Equal(CONTENT))
			})

			AfterEach(func() {
				note := noteService.FindAll()[0]
				noteService.Delete(note)
			})
		})
		Context("If there are no notes in the database", func() {
			It("should return an empty list", func() {
				notes := noteService.FindAll()
				Ω(notes).Should(BeEmpty())
			})
		})
	})
})

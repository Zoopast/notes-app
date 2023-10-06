package service

import (
	"github.com/Zoopast/notes-app/model"
	"github.com/Zoopast/notes-app/repository"
)

type NoteService interface {
	Save(model.Note) error
	Update(model.Note) error
	Delete(model.Note) error
	FindAll() []model.Note
	Find(id uint64) model.Note
}

type noteService struct {
	repository repository.NoteRepository
}

func New(noteRepository repository.NoteRepository) NoteService {
	return &noteService{
		repository: noteRepository,
	}
}
func (service *noteService) Save(note model.Note) error {
	service.repository.Save(note)
	return nil
}
func (service *noteService) Update(note model.Note) error {
	service.repository.Update(note)
	return nil
}
func (service *noteService) Delete(note model.Note) error {
	service.repository.Delete(note)
	return nil
}
func (service *noteService) FindAll() []model.Note {
	return service.repository.FindAll()
}

func (service *noteService) Find(id uint64) model.Note {
	return service.repository.Find(id)
}

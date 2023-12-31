package repository

import (
	"github.com/Zoopast/notes-app/initializer"
	"github.com/Zoopast/notes-app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NoteRepository interface {
	Save(model.Note)
	Update(model.Note)
	Delete(model.Note)
	FindAll() []model.Note
	Find(id uint64) model.Note
}

type database struct {
	connection *gorm.DB
}

func NewNoteRepository(configPath string) NoteRepository {
	config := initializer.LoadConfig(configPath)

	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&model.Note{})

	return &database{
		connection: db,
	}
}

func (db *database) Save(note model.Note) {
	db.connection.Create(&note)
}
func (db *database) Update(note model.Note) {
	db.connection.Save(&note)
}
func (db *database) Delete(note model.Note) {
	db.connection.Delete(&note)
}
func (db *database) FindAll() []model.Note {
	var notes []model.Note

	db.connection.Set("gorm:auto_preload", true).Find(&notes)
	return notes
}

func (db *database) Find(id uint64) model.Note {
	var note model.Note
	db.connection.First(&note, id)

	return note
}

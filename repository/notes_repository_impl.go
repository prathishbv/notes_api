package repository

import (
	"errors"
	"fmt"

	"github.com/prathishbv/notes-api/data/request"
	"github.com/prathishbv/notes-api/helper"
	"github.com/prathishbv/notes-api/model"
	"gorm.io/gorm"
)

type NotesRepositoryImpl struct {
	Db *gorm.DB
}

func NewNotesRepositoryImpl(Db *gorm.DB) NotesRepository {
	return &NotesRepositoryImpl{Db: Db}
}

func (r *NotesRepositoryImpl) GetNotes(userID int) []model.Note {
	var notes []model.Note
	result := r.Db.Where("created_by_id = ?", userID).Find(&notes)
	helper.ErrorPanic(result.Error)
	return notes
}

func (r *NotesRepositoryImpl) GetNoteByID(noteID int, userID int) (model.Note, error) {
	var note model.Note
	result := r.Db.First(&note, "id = ? AND created_by_id = ?", noteID, userID)
	if result.Error != nil {
		return model.Note{}, errors.New("note not found")
	}
	return note, nil
}

func (r *NotesRepositoryImpl) CreateNote(note model.Note, userID int) int {
    result := r.Db.Create(&note)
    helper.ErrorPanic(result.Error)

    return note.ID
}


func (r *NotesRepositoryImpl) UpdateNote(noteID int, request request.UpdateNoteRequest, userID int) error {
    // Fetch the existing note from the database
    existingNote, err := r.GetNoteByID(noteID, userID)
    if err != nil {
        return err
    }

    if existingNote.CreatedByID != userID {
        return fmt.Errorf("user does not own the note with ID %d", noteID)
    }

    existingNote.Title = request.Title
    existingNote.Content = request.Content

    result := r.Db.Save(&existingNote)
    if result.Error != nil {
        return result.Error
    }

    return nil
}


func (r *NotesRepositoryImpl) DeleteNote(noteID int, userID int) error {
    
    note := model.Note{ID: noteID, CreatedByID: userID}

    result := r.Db.Delete(&note)
    if result.Error != nil {
        helper.ErrorPanic(result.Error)
        return result.Error
    }

    return nil
}

func (r *NotesRepositoryImpl) CreateSharedNote(sharedNote model.SharedNote) {
	result := r.Db.Create(&sharedNote)
	helper.ErrorPanic(result.Error)

}

func (r *NotesRepositoryImpl) FindUserByUsername(username string) (model.Users, error) {
	var user model.Users
	result := r.Db.First(&user, "username = ?", username)
	if result.Error != nil {
		return model.Users{}, errors.New("user not found")
	}
	return user, nil
}

func (r *NotesRepositoryImpl) IsNoteSharedWithUser(noteID int, userID int) bool {
	var sharedNote model.SharedNote
	result := r.Db.First(&sharedNote, "note_id = ? AND user_id = ?", noteID, userID)
	return result.Error == nil
}

func (r *NotesRepositoryImpl) SearchNotes(query string, userID int) []model.Note {
	var notes []model.Note
	result := r.Db.Where("created_by_id = ? AND (title ILIKE ? OR content ILIKE ?)", userID, "%"+query+"%", "%"+query+"%").Find(&notes)
	helper.ErrorPanic(result.Error)
	return notes
}


func (r *NotesRepositoryImpl) ShareNote(noteID int, request request.ShareNoteRequest, userID int) error {

    note, err := r.GetNoteByID(noteID, userID)
    if err != nil {
        return err 
    }


    if note.CreatedByID != userID {
        return fmt.Errorf("user does not own the note with ID %d", noteID)
    }

    sharedNote := model.SharedNote{
        NoteID:     noteID,
        SharedByID: request.SharedWith,
    }

    r.CreateSharedNote(sharedNote)
    

    return nil
}

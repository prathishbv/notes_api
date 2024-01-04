package repository

import (
	"github.com/prathishbv/notes-api/data/request"
	"github.com/prathishbv/notes-api/model"
)

type NotesRepository interface {
	GetNotes(userID int) []model.Note
	GetNoteByID(noteID int, userID int) (model.Note, error)
	CreateNote(note model.Note, userID int) int
	UpdateNote(noteID int, request request.UpdateNoteRequest, userID int) error
	DeleteNote(noteID int, userID int) error
	CreateSharedNote(sharedNote model.SharedNote)
	FindUserByUsername(username string) (model.Users, error)
	IsNoteSharedWithUser(noteID int, userID int) bool
	SearchNotes(query string, userID int) []model.Note
	ShareNote(noteID int, request request.ShareNoteRequest, userID int) error

}
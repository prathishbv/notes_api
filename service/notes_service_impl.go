package service

import (
    "errors"
    "time"

    "github.com/go-playground/validator/v10"
    "github.com/prathishbv/notes-api/data/request"
    "github.com/prathishbv/notes-api/helper"
    "github.com/prathishbv/notes-api/model"
    "github.com/prathishbv/notes-api/repository"
)

type NotesServiceImpl struct {
    NotesRepository repository.NotesRepository
    Validate        *validator.Validate
}

func NewNotesServiceImpl(notesRepository repository.NotesRepository, validate *validator.Validate) NotesService {
    return &NotesServiceImpl{
        NotesRepository: notesRepository,
        Validate:        validate,
    }
}

func (s *NotesServiceImpl) GetNotes(userID int) []model.Note {
    return s.NotesRepository.GetNotes(userID)
}

func (s *NotesServiceImpl) GetNoteByID(noteID int, userID int) (model.Note, error) {
    note, err := s.NotesRepository.GetNoteByID(noteID, userID)
    if err != nil {
        return model.Note{}, errors.New("note not found")
    }
    return note, nil
}

func (s *NotesServiceImpl) CreateNote(note model.Note, userID int) int {
    note.CreatedByID = userID

    note.CreatedAt = time.Now()
    note.UpdatedAt = time.Now()

    err := s.Validate.Struct(note)
    helper.ErrorPanic(err)

    return s.NotesRepository.CreateNote(note, userID)
}

func (s *NotesServiceImpl) UpdateNote(noteID int, request request.UpdateNoteRequest, userID int) error {
    note, err := s.NotesRepository.GetNoteByID(noteID, userID)
    if err != nil {
        return errors.New("note not found")
    }

    note.Title = request.Title
    note.Content = request.Content
    note.UpdatedAt = time.Now()

    err = s.NotesRepository.UpdateNote(noteID, request, userID)
    if err != nil {
        return err
    }

    return nil
}

func (s *NotesServiceImpl) DeleteNote(noteID int, userID int) error {
    _, err := s.NotesRepository.GetNoteByID(noteID, userID)
    if err != nil {
        return err 
    }

    err = s.NotesRepository.DeleteNote(noteID, userID)
    if err != nil {
        return err
    }

    return nil
}

func (s *NotesServiceImpl) ShareNote(noteID int, request request.ShareNoteRequest, userID int) error {
    note, err := s.NotesRepository.GetNoteByID(noteID, userID)
    if err != nil {
        return errors.New("note not found")
    }

    // Check if the user to be shared with exists
    sharedWithUser, err := s.NotesRepository.FindUserByUsername(request.Username)
    if err != nil {
        return errors.New("user not found")
    }

    // Check if the note is already shared with the user
    if s.NotesRepository.IsNoteSharedWithUser(noteID, sharedWithUser.Id) {
        return errors.New("note already shared with the user")
    }

    // Create a shared note entry
    sharedNote := model.SharedNote{
        NoteID:     note.ID,
        UserID:     sharedWithUser.Id,
        SharedByID: userID,
        CreatedAt:  time.Now(),
    }

    s.NotesRepository.CreateSharedNote(sharedNote)
    return nil
}

func (s *NotesServiceImpl) SearchNotes(query string, userID int) []model.Note {
    return s.NotesRepository.SearchNotes(query, userID)
}

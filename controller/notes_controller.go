package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prathishbv/notes-api/data/request"
	"github.com/prathishbv/notes-api/data/response"
	"github.com/prathishbv/notes-api/helper"
	"github.com/prathishbv/notes-api/model"
	"github.com/prathishbv/notes-api/service"
)

type NotesController struct {
    notesService service.NotesService
}

func NewNotesController(service service.NotesService) *NotesController {
    return &NotesController{notesService: service}
}

func (controller *NotesController) GetNotes(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.Users)
	notes := controller.notesService.GetNotes(currentUser.Id)
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetch all notes!",
		Data:    notes,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotesController) GetNote(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.Users)
	noteID := helper.GetIDParam(ctx, "id")
	note, err := controller.notesService.GetNoteByID(noteID, currentUser.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.NewErrorResponse(http.StatusNotFound, "Note not found"))
		return
	}
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetch note!",
		Data:    note,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func convertToNote(createNoteRequest request.CreateNoteRequest, userID int) model.Note {
    return model.Note{
        Title:       createNoteRequest.Title,
        Content:     createNoteRequest.Content,
        CreatedByID: userID,
    }
}

func (controller *NotesController) CreateNote(ctx *gin.Context) {
    currentUser := ctx.MustGet("currentUser").(model.Users)
    createNoteRequest := request.CreateNoteRequest{}
    err := ctx.ShouldBindJSON(&createNoteRequest)
    helper.ErrorPanic(err)

    // Convert the request to a model.Note
    note := convertToNote(createNoteRequest, currentUser.Id)

    // Call the service method with the converted note
    noteID := controller.notesService.CreateNote(note, currentUser.Id)

    webResponse := response.Response{
        Code:    http.StatusCreated,
        Status:  "Created",
        Message: "Successfully created note!",
        Data:    map[string]int{"noteID": noteID},
    }
    ctx.JSON(http.StatusCreated, webResponse)
}

// func (controller *NotesController) CreateNote(ctx *gin.Context) {
// 	currentUser := ctx.MustGet("currentUser").(model.Users)
// 	createNoteRequest := request.CreateNoteRequest{}
// 	err := ctx.ShouldBindJSON(&createNoteRequest)
// 	helper.ErrorPanic(err)

// 	noteID := controller.notesService.CreateNote(createNoteRequest, currentUser.Id)

// 	webResponse := response.Response{
// 		Code:    http.StatusCreated,
// 		Status:  "Created",
// 		Message: "Successfully created note!",
// 		Data:    map[string]int{"noteID": noteID},
// 	}
// 	ctx.JSON(http.StatusCreated, webResponse)
// }

func (controller *NotesController) UpdateNote(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.Users)
	noteID := helper.GetIDParam(ctx, "id")

	updateNoteRequest := request.UpdateNoteRequest{}
	err := ctx.ShouldBindJSON(&updateNoteRequest)
	helper.ErrorPanic(err)

	err = controller.notesService.UpdateNote(noteID, updateNoteRequest, currentUser.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.NewErrorResponse(http.StatusNotFound, "Note not found"))
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully updated note!",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotesController) DeleteNote(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.Users)
	noteID := helper.GetIDParam(ctx, "id")

	err := controller.notesService.DeleteNote(noteID, currentUser.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.NewErrorResponse(http.StatusNotFound, "Note not found"))
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully deleted note!",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotesController) ShareNote(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.Users)
	noteID := helper.GetIDParam(ctx, "id")

	shareNoteRequest := request.ShareNoteRequest{}
	err := ctx.ShouldBindJSON(&shareNoteRequest)
	helper.ErrorPanic(err)

	err = controller.notesService.ShareNote(noteID, shareNoteRequest, currentUser.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, response.NewErrorResponse(http.StatusNotFound, "Note not found"))
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully shared note!",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *NotesController) SearchNotes(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(model.Users)
	query := ctx.Query("q")

	notes := controller.notesService.SearchNotes(query, currentUser.Id)

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok",
		Message: "Successfully fetched notes based on search!",
		Data:    notes,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

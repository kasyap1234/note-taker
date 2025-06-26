package api

import (
	"net/http"
	"note-taker/internal"

	"github.com/labstack/echo/v4"
)

// NoteHandler is the presentation layer, responsible for handling HTTP requests.
// It uses the NoteService to perform business operations.
type NoteHandler struct {
	service *internal.NoteService
}

// NewNoteHandler creates a new NoteHandler.
func NewNoteHandler(service *internal.NoteService) *NoteHandler {
	return &NoteHandler{service: service}
}

// CreateNoteRequest defines the structure for the incoming JSON payload.
type CreateNoteRequest struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
}

// CreateNote is an Echo handler for creating a new note.
func (h *NoteHandler) CreateNote(c echo.Context) error {
	req := new(CreateNoteRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	params := internal.CreateNoteParams{
		Title:   req.Title,
		Content: req.Content,
	}

	note, err := h.service.CreateNote(c.Request().Context(), params)
	if err != nil {
		// The service layer provides a clear error, which we can return to the user.
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, note)
}

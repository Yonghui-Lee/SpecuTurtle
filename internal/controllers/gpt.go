package controllers

import (
	"encoding/json"
	"net/http"
	"satellity/internal/session"
	"satellity/internal/views"

	"github.com/dimfeld/httptreemux"
)

type gptImpl struct{}

type gptRequest struct {
	Content     string `json:"content"`
}

func registerGpt(router *httptreemux.Group) {
	impl := &gptImpl{}

	router.POST("/gpt", impl.index)
}

func (impl *gptImpl) index(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	var body gptRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	views.RenderGpt(w, r, body.Content)
}

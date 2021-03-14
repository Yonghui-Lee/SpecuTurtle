package views

import (
	"net/http"
)

type gptView struct {
	Type     string `json:"type"`
	Response string `json:"response"`
}

func buildGpt(response string) gptView {
	return gptView{
		Type:           "gpt_response",
		Response:       response,
	}
}

func RenderGpt(w http.ResponseWriter, r *http.Request, content string) {
	RenderResponse(w, r, buildGpt(content))
}

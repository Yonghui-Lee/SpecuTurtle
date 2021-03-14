package controllers

import (
	"encoding/json"
	"net/http"
	"satellity/internal/session"
	"satellity/internal/views"
	"satellity/internal/configs"
	"bytes"
	"io/ioutil"

	"github.com/dimfeld/httptreemux"
)

type gptImpl struct{}

type gptRequest struct {
	Content     string `json:"content"`
}

type chatResponse struct {
	Result      string `json:"result"`
	Status      bool `json:"status"`
}

type poetryResponse struct {
	Result      string `json:"result"`
	Status      bool `json:"status"`
}

type qaResponse struct {
	Result      qaResult `json:"result"`
	Status      bool `json:"status"`
}

type qaResult struct {
    Generation string  `json:"content"`
    Type 	   int 	   `json:"type"`
    Bullet_id  string  `json:"bullet_id"`
}


func registerGpt(router *httptreemux.Group) {
	impl := &gptImpl{}

	router.POST("/gpt/:method", impl.index)
}

func (impl *gptImpl) index(w http.ResponseWriter, r *http.Request, params map[string]string) {
	var body gptRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
		return
	}
	
	
	info := make(map[string]string)    
	info["token"] = configs.AppConfig.GptToken
	info["app"] = params["method"]
	info["content"] = body.Content
 
	bytesData, err := json.Marshal(info)
	if err != nil {
		views.RenderErrorResponse(w, r, err)
		return
	}
 
	reader := bytes.NewReader(bytesData)
 
	url := "http://lab.aminer.cn/isoa-2021/gpt" 
	request, err := http.NewRequest("POST", url, reader)
	defer request.Body.Close()  
	if err != nil {
		views.RenderErrorResponse(w, r, err)
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")   
	
	client := http.Client{}
	resp, err := client.Do(request) 
	if err != nil {
		views.RenderErrorResponse(w, r, err)
		return
	}
 
	respBytes, err := ioutil.ReadAll(resp.Body)   
	if err != nil {
		views.RenderErrorResponse(w, r, err)
		return
	}

	switch params["method"] {
    case "qa":
        var resBody qaResponse
	
		err = json.Unmarshal(respBytes, &resBody)
		if err != nil {
			views.RenderErrorResponse(w, r, err)
			return
		}

		views.RenderGpt(w, r, resBody.Result.Generation)
    case "chat":
        var resBody chatResponse
	
		err = json.Unmarshal(respBytes, &resBody)
		if err != nil {
			views.RenderErrorResponse(w, r, err)
			return
		}

		views.RenderGpt(w, r, resBody.Result)
	case "poetry":
        var resBody poetryResponse
	
		err = json.Unmarshal(respBytes, &resBody)
		if err != nil {
			views.RenderErrorResponse(w, r, err)
			return
		}

		views.RenderGpt(w, r, resBody.Result)
    default:
        views.RenderErrorResponse(w, r, session.BadRequestError(r.Context()))
	}

	
}

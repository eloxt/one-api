package vertexai

import (
	"net/http"

	"github.com/eloxt/one-api/common/ctxkey"
	"github.com/eloxt/one-api/relay/adaptor/gemini"
	"github.com/eloxt/one-api/relay/adaptor/openai"
	"github.com/eloxt/one-api/relay/relaymode"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/eloxt/one-api/relay/meta"
	"github.com/eloxt/one-api/relay/model"
)

var ModelList = []string{
	"gemini-pro", "gemini-pro-vision",
	"gemini-1.5-pro-001", "gemini-1.5-flash-001",
	"gemini-1.5-pro-002", "gemini-1.5-flash-002",
	"gemini-2.0-flash-exp", "gemini-2.0-flash-thinking-exp",
}

type Adaptor struct {
}

func (a *Adaptor) ConvertRequest(c *gin.Context, relayMode int, request *model.GeneralOpenAIRequest) (any, error) {
	if request == nil {
		return nil, errors.New("request is nil")
	}

	geminiRequest := gemini.ConvertRequest(*request)
	c.Set(ctxkey.RequestModel, request.Model)
	c.Set(ctxkey.ConvertedRequest, geminiRequest)
	return geminiRequest, nil
}

func (a *Adaptor) DoResponse(c *gin.Context, resp *http.Response, meta *meta.Meta) (usage *model.Usage, err *model.ErrorWithStatusCode, id string) {
	if meta.IsStream {
		var responseText string
		err, responseText = gemini.StreamHandler(c, resp)
		usage = openai.ResponseText2Usage(responseText, meta.ActualModelName, meta.PromptTokens)
	} else {
		switch meta.Mode {
		case relaymode.Embeddings:
			err, usage = gemini.EmbeddingHandler(c, resp)
		default:
			err, usage = gemini.Handler(c, resp, meta.PromptTokens, meta.ActualModelName)
		}
	}
	return
}

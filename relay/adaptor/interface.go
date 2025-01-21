package adaptor

import (
	"io"
	"net/http"

	"github.com/eloxt/one-api/relay/meta"
	"github.com/eloxt/one-api/relay/model"
	"github.com/gin-gonic/gin"
)

type Adaptor interface {
	Init(meta *meta.Meta)
	GetRequestURL(meta *meta.Meta) (string, error)
	SetupRequestHeader(c *gin.Context, req *http.Request, meta *meta.Meta) error
	ConvertRequest(c *gin.Context, relayMode int, request *model.GeneralOpenAIRequest) (any, error)
	ConvertImageRequest(request *model.ImageRequest) (any, error)
	DoRequest(c *gin.Context, meta *meta.Meta, requestBody io.Reader) (*http.Response, error)
	DoResponse(c *gin.Context, resp *http.Response, meta *meta.Meta) (usage *model.Usage, err *model.ErrorWithStatusCode, id string)
	GetModelList() []string
	GetChannelName() string
}

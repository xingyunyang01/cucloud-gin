package middlewares

import "github.com/gin-gonic/gin"

type AddVersion struct {
}

func NewAddVersion() *AddVersion {
	return &AddVersion{}
}

func (this *AddVersion) OnRequest(ctx *gin.Context) error {
	return nil
}

func (this *AddVersion) OnResponse(result interface{}) (interface{}, error) {
	if str, ok := result.(string); ok {
		str = str + "_version_1.0"
		return str, nil
	}
	return result, nil
}

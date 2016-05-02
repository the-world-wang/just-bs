package public
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/dto"
	"just.com/service/image"
	"net/http"
	"just.com/common"
)

func UploadCallbackHandle(c *gin.Context) {
	context := action.GetContext(c)
	request := new(dto.UploadCallbackRequest)
	bindErr := c.Bind(request)
	if bindErr != nil {
		context.Log.Println(bindErr)
		return
	}
	context.Log.Println(request)
	if common.IsEmpty(request.Key) {
		c.JSON(http.StatusOK, nil)
		return
	}
	imageService := new(image.ImageService)
	imageService.Add("", request.Key, "", request.W, request.H)
}

package course
import (
	"github.com/gin-gonic/gin"
	"just.com/action"
	"just.com/service/course"
	"just.com/common"
	"just.com/err"
)

func CourseDeleteHandle(c *gin.Context) {
	context := action.GetContext(c)
	courseId := c.Param("course_id")
	if common.IsEmpty(courseId) {
		context.Response.Error = err.NO_COURSE_ID_FOUND
	}
	courseService := service.NewCourseService(context.Session, context.Log)
	deleteErr := courseService.Delete(courseId)
	if deleteErr != nil {
		context.Log.Println(deleteErr)
		context.Response.Error = deleteErr
		return
	}
}

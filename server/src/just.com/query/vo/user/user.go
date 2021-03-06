package user
import (
	"just.com/query/vo/file"
	"github.com/go-xorm/xorm"
	"log"
	"just.com/model/db/table"
)

type UserVo struct {
	UUID     string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Icon     *file.ImageVo `json:"icon"`
	RoleName string `json:"role_name"`
}

func LoadUserVo(userId string, session *xorm.Session, log *log.Logger) *UserVo {
	userTable := new(table.UserTable)
	getFlag, getErr := session.Id(userId).Get(userTable)
	if getFlag == false {
		if getErr != nil {
			log.Println(getErr)
		}
		return nil
	}
	return LoadUserVoByTable(userTable)
}

func LoadUserVoByTable(table *table.UserTable) *UserVo {
	uv := new(UserVo)
	uv.UUID = table.UUID
	uv.Name = table.Name
	uv.Email = table.Email
	uv.RoleName = table.RoleName
	uv.Icon = file.NewImageVo(table.IconUrl, table.IconWidth, table.IconHeight)
	return uv
}
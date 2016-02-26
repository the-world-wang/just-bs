package model
import "time"

type ImageTable struct {
	UUID       string    `xorm:"pk 'UUID'"`
	Name       string    `xorm:"'NAME'"`
	Url        string    `xorm:"'URL'"`
	Width      int64    `xorm:"'WIDTH'"`
	Height     int64    `xorm:"'HEGITH'"`
	CreateTime time.Time    `xorm:"created 'CREATED_TIME'"`
}

type FileTable struct {
	UUID         string `xorm:"pk 'UUID'"`
	Name         string `xorm:"'NAME'"`
	URL          string `xorm:"'URL'"`
	CourseID     string `xorm:"'COURSE_ID'"`
	CreateTime   time.Time `xorm:"created 'CREATE_TIME'"`
	CreateUser   string `xorm:"'CREATE_USER'"`
	UpdateTime   time.Time `xorm:"updated 'UPDATE_TIME'"`
	UpdateUser   string `xorm:"'UPDATE_USER'"`
	FrozenStatus string `xorm:"'FROZEN_STATUS'"`
	FrozenTime   time.Time `xorm:"deleted 'FROZEN_TIME'"`
	Version      int64 `xorm:"version 'VERSION'"`
}

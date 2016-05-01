package service
import "errors"

const (
	SERVICE_TOKEN_CHECK_ERR = "SERVICE_TOKEN_CHECK_ERR"
	SERVICE_TOKEN_MAKE_ERR = "SERVICE_TOKEN_MAKE_ERR"

	SERVICE_COURSE_UPDATE_ERR = "SERVICE_COURSE_UPDATE_ERR"
	SERVICE_COURSE_DELETE_ERR = "SERVICE_COURSE_DELETE_ERR"
// mark
	SERVICE_COURSE_MARK_ERR = "SERVICE_COURSE_MARK_ERR"
	SERVICE_COURSE_MARK_CANCEL_ERR = "SERVICE_COURSE_MARK_CANCEL_ERR"
// flush
	SERVICE_COURSE_FLUSH_MARK_NUM_ERR = "SERVICE_COURSE_FLUSH_MARK_NUM_ERR"
	SERVICE_COURSE_FLUSH_COMMENT_NUM_ERR = "SERVICE_COURSE_FLUSH_COMMENT_NUM_ERR"
	SERVICE_COURSE_FLUSH_POINT_ERR = "SERVICE_COURSE_FLUSH_POINT_ERR"
// comment
	SERVICE_COURSE_COMMENT_ADD_ERR = "SERVICE_COURSE_COMMENT_ADD_ERR"
	SERVICE_COURSE_COMMENT_DELETE_ERR = "SERVICE_COURSE_COMMENT_DELETE_ERR"
// point
	SERVICE_COURSE_POINT_ADD_ERR = "SERVICE_COURSE_POINT_ADD_ERR"
	SERVICE_COURSE_POINT_UPDATE_ERR = "SERVICE_COURSE_POINT_UPDATE_ERR"

	SERVICE_FILE_ADD_ERR = "SERVICE_FILE_ADD_ERR"
	SERVICE_FILE_UPDATE_ERR = "SERVICE_FILE_UPDATE_ERR"
	SERVICE_FILE_DELETE_ERR = "SERVICE_FILE_DELETE_ERR"

	SERVICE_IMAGE_FIND_BY_ID_ERR = "SERVICE_IMAGE_FIND_BY_ID_ERR"



	SERVICE_RBAC_LOAD_ERR = "SERVICE_RBAC_LOAD_ERR"
	SERVICE_RBAC_ROLE_ASSIGN = "SERVICE_RBAC_ROLE_ASSIGN"
	SERVICE_RBAC_ROLE_REVOKE = "SERVICE_RBAC_ROLE_REVOKE"
	SERVICE_RBAC_ROLE_ADD_ERR = "SERVICE_RBAC_ROLE_ADD_ERR"
	SERVICE_RBAC_ROLE_FIND_BY_NAME_ERR = "SERVICE_RBAC_ROLE_FIND_BY_NAME_ERR"
	SERVICE_RBAC_ROLE_ADD_PERMISSION_ERR = "SERVICE_RBAC_ROLE_ADD_PERMISSION_ERR"
	SERVICE_RBAC_PERMISSION_ADD_ERR = "SERVICE_RBAC_PERMISSION_ADD_ERR"

// user
	SERVICE_USER_ADD_ERR = "SERVICE_USER_ADD_ERR"
)

var (
	SERVICE_USER_REGISTER_ERR = errors.New("SERVICE_USER_REGISTER_ERR")
	SERVICE_USER_UPDATE_ERR = errors.New("SERVICE_USER_UPDATE_ERR")
	SERVICE_USER_REGISTER_EMAIL_EXISTS_ERR = errors.New("SERVICE_USER_REGISTER_EMAIL_EXISTS_ERR")
	SERVICE_USER_ACTIVE_ERR = errors.New("SERVICE_USER_ACTIVE_ERR")
	SERVICE_IMAGE_ADD_ERR = errors.New("SERVICE_IMAGE_ADD_ERR")
	SERVICE_COURSE_ADD_ERR = errors.New("SERVICE_COURSE_ADD_ERR")
)
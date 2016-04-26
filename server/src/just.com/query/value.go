package query
import "errors"

var(
	QUERY_COURSE_LOAD_ERR = errors.New("QUERY_COURSE_LOAD_ERR")
	QUERY_COURSE_LOAD_FROM_TABLE_ERR = errors.New("QUERY_COURSE_LOAD_FROM_TABLE_ERR")
	QUERY_COURSE_LOAD_LIST_ERR = errors.New("QUERY_COURSE_LOAD_LIST_ERR")
)

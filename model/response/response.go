package response

// type ErrorcodeEnum
type ShowTypeEnum int8

const (
	SILENT       ShowTypeEnum = 0
	WARM         ShowTypeEnum = 1
	ERROR        ShowTypeEnum = 2
	NOTIFICATION ShowTypeEnum = 4
	PAGE         ShowTypeEnum = 9
)

type Response struct {
	Success      bool         `json:"success"`
	Data         interface{}  `json:"data"`
	Total        int64        `json:"total"`
	ErrorMessage string       `json:"errorMessage"`
	ErrorCode    int32        `json:"errorCode"`
	ShowType     ShowTypeEnum `json:"showType"`
}

package biz

import "net/http"

const (
	// 预定义, 通常情况下, 0 => 表示成功, 其余表示错误或者异常行为
	BizSuccess = 0
	BizFailure = -1

	// 400 +
	BizUnauthorized    = http.StatusUnauthorized
	BizForbidden       = http.StatusForbidden
	BizRequestTimeout  = http.StatusRequestTimeout
	BizMehodNotAllowed = http.StatusMethodNotAllowed

	// 500 +
	BizInternalServerError = http.StatusInternalServerError
	BizNotImplemented      = http.StatusNotImplemented
	BizBadGateway          = http.StatusBadGateway

	// 错误化状态码：< -1000
	BizUserNotFound = -1001
	BizTokenExpired = -2001

	// 业务化状态码: >1000
	BizUserNotHasAdminRole = 1001
)

var bizMessage = map[int]string{
	BizSuccess: "",
	BizFailure: "",

	BizUnauthorized:    http.StatusText(BizUnauthorized),
	BizForbidden:       http.StatusText(BizForbidden),
	BizRequestTimeout:  http.StatusText(BizRequestTimeout),
	BizMehodNotAllowed: http.StatusText(BizMehodNotAllowed),

	BizInternalServerError: http.StatusText(BizInternalServerError),
	BizNotImplemented:      http.StatusText(BizNotImplemented),
	BizBadGateway:          http.StatusText(BizBadGateway),

	BizUserNotFound: "User not found",
	BizTokenExpired: "Token expired",

	BizUserNotHasAdminRole: "User does not have admin role",
}

func BizMessage(code int) *string {

	text := bizMessage[code]

	if text == "" {
		return nil
	}

	return &text
}

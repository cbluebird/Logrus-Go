package apiException

import "net/http"

type Error struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}

var (
	ServerError           = NewError(http.StatusInternalServerError, 200500, "系统异常，请稍后重试!")
	OpenIDError           = NewError(http.StatusInternalServerError, 200500, "openid错误")
	ParamError            = NewError(http.StatusInternalServerError, 200501, "参数错误")
	NotAdmin              = NewError(http.StatusBadRequest, 200502, "该用户无管理员权限")
	UserNotFind           = NewError(http.StatusBadRequest, 200503, "该用户不存在")
	NotLogin              = NewError(http.StatusBadRequest, 200503, "未登录")
	NoThatPasswordOrWrong = NewError(http.StatusBadRequest, 200504, "密码错误")
	HttpTimeout           = NewError(http.StatusInternalServerError, 200505, "网络连接超时，请稍后重试!")
	RequestError          = NewError(http.StatusInternalServerError, 200506, "服务响应错误，请稍后重试!")
	NotBindYxy            = NewError(http.StatusBadRequest, 200507, "该手机号或用户未绑定或未注册易校园")
	UserAlreadyExisted    = NewError(http.StatusBadRequest, 200508, "该用户已激活")
	WrongVerificationCode = NewError(http.StatusInternalServerError, 200509, "验证码错误")
	StudentNumAndIidError = NewError(http.StatusBadRequest, 200510, "该学号或身份证不存在或者不匹配，请重新输入")
	PwdError              = NewError(http.StatusBadRequest, 200511, "密码长度必须在6~20位之间")
	ReactiveError         = NewError(http.StatusBadRequest, 200512, "该通行证已经存在，请重新输入")
	StudentIdError        = NewError(http.StatusBadRequest, 200513, "学号格式不正确，请重新输入")
	YxySessionExpired     = NewError(http.StatusBadRequest, 200514, "一卡通登陆过期，请稍后再试")
	YxyNeedCaptcha        = NewError(http.StatusBadRequest, 500515, "请输入验证码")
	WrongCaptcha          = NewError(http.StatusBadRequest, 200516, "图形验证码错误")
	WrongPhoneNum         = NewError(http.StatusBadRequest, 200517, "手机号格式不正确")
	ImgTypeError          = NewError(http.StatusBadRequest, 200518, "图片类型有误")
	NotInit               = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	NotFound              = NewError(http.StatusNotFound, 200404, http.StatusText(http.StatusNotFound))
	Unknown               = NewError(http.StatusInternalServerError, 300500, "系统未知异常，请稍后重试!")
)

func OtherError(message string) *Error {
	return NewError(http.StatusForbidden, 100403, message)
}

func (e *Error) Error() string {
	return e.Msg
}

func NewError(statusCode, Code int, msg string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       Code,
		Msg:        msg,
	}
}

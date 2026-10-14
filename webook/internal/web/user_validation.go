package web

const (
	// 邮箱：常见邮箱格式，本地部分最长 64 位，域名至少包含一级后缀。
	emailRegexPattern = `^[A-Za-z0-9][A-Za-z0-9._%+-]{0,63}@[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?(?:\.[A-Za-z0-9](?:[A-Za-z0-9-]{0,61}[A-Za-z0-9])?)+$`

	// 密码：8-32 位，只允许字母、数字和以下特殊字符。
	passwordRegexPattern        = `^[A-Za-z0-9!@#$%^&*]{8,32}$`
	passwordLetterRegexPattern  = `[A-Za-z]`
	passwordDigitRegexPattern   = `[0-9]`
	passwordSpecialRegexPattern = `[!@#$%^&*]`
)

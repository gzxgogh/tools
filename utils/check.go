package utils

import (
	"regexp"
	"strings"
)

const (
	// 中国大陆手机号码正则匹配, 不是那么太精细
	// 只要是 13,14,15,17,18,19 开头的 11 位数字就认为是中国手机号
	chinaMobilePattern = `^1[0-9]{10}$`
	// 用户昵称的正则匹配, 合法的字符有 0-9, A-Z, a-z, _, 汉字
	// 字符 '_' 只能出现在中间且不能重复, 如 "__"
	nicknamePattern = `^[a-z0-9A-Z\p{Han}]+(_[a-z0-9A-Z\p{Han}]+)*?$`
	// 用户名的正则匹配, 合法的字符有 0-9, A-Z, a-z, _
	// 第一个字母不能为 _, 0-9
	// 最后一个字母不能为 _, 且 _ 不能连续
	usernamePattern = `^[a-zA-Z][a-z0-9A-Z]*(_[a-z0-9A-Z]+)*?$`
	// 电子邮箱的正则匹配, 考虑到各个网站的 mail 要求不一样, 这里匹配比较宽松
	// 邮箱用户名可以包含 0-9, A-Z, a-z, -, _, .
	// 开头字母不能是 -, _, .
	// 结尾字母不能是 -, _, .
	// -, _, . 这三个连接字母任意两个不能连续, 如不能出现 --, __, .., -_, -., _.
	// 邮箱的域名可以包含 0-9, A-Z, a-z, -
	// 连接字符 - 只能出现在中间, 不能连续, 如不能 --
	// 支持多级域名, x@y.z, x@y.z.w, x@x.y.z.w.e
	mailPattern = `^[a-z0-9A-Z]+([_\-\.][a-z0-9A-Z]+)*?@[a-z0-9A-Z]+([\-\.][a-z0-9A-Z]+)*?\.[a-zA-Z]{2,}$`

	chineseNamePattern   = "^\\p{Han}+(\u00B7\\p{Han}+)*?$"
	chineseNameExPattern = "^\\p{Han}+([\u00B7\u2022\u2027\u30FB\u002E\u0387\u16EB\u2219\u22C5\uFF65\u05BC]\\p{Han}+)*?$"
)

var (
	chinaMobileRegexp   = regexp.MustCompile(chinaMobilePattern)
	nicknameRegexp      = regexp.MustCompile(nicknamePattern)
	usernameRegexp      = regexp.MustCompile(usernamePattern)
	mailRegexp          = regexp.MustCompile(mailPattern)
	chineseNameRegexp   = regexp.MustCompile(chineseNamePattern)
	chineseNameExRegexp = regexp.MustCompile(chineseNameExPattern)
)

//检验是否为合法的中国手机号, 不是那么太精细,只要是 13,14,15,18 开头的 11 位数字就认为是中国手机号
func IsChinaMobile(s string) bool {
	if len(s) != 11 {
		return false
	}
	return chinaMobileRegexp.MatchString(s)
}

// 检验是否为合法的昵称, 合法的字符有 0-9, A-Z, a-z, _, 汉字,字符 '_' 只能出现在中间且不能重复, 如 "__"
func IsNickname(s string) bool {
	if len(s) == 0 {
		return false
	}
	return nicknameRegexp.MatchString(s)
}

// 检验是否为合法的用户名, 合法的字符有 0-9, A-Z, a-z, _
// 第一个字母不能为 _, 0-9
// 最后一个字母不能为 _, 且 _ 不能连续
func IsUserName(s string) bool {
	if len(s) == 0 {
		return false
	}
	return usernameRegexp.MatchString(s)
}

// 检验是否为合法的电子邮箱, 考虑到各个网站的 mail 要求不一样, 这里匹配比较宽松
// 邮箱用户名可以包含 0-9, A-Z, a-z, -, _, .
// 开头字母不能是 -, _, .
// 结尾字母不能是 -, _, .
// -, _, . 这三个连接字母任意两个不能连续, 如不能出现 --, __, .., -_, -., _.
// 邮箱的域名可以包含 0-9, A-Z, a-z, -
// 连接字符 - 只能出现在中间, 不能连续, 如不能 --
// 支持多级域名, x@y.z, x@y.z.w, x@x.y.z.w.e
func IsMail(s string) bool {
	if len(s) < 6 { // x@x.xx
		return false
	}
	return mailRegexp.MatchString(s)
}

// IsChineseName 检验是否为有效的中文姓名(比如 张三, 李四, 张三·李四)
func IsChineseName(s string) bool {
	return chineseNameRegexp.MatchString(s)
}

//判断是否是18或15位身份证
func IsIdCard(cardNo string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
		return false
	}
	return true
}

// CheckSqlValidate 检查是否含有可能产生注入的非法字符
// 返回值为true时表示含有非法字符，同时返回的字符串值为匹配到的非法字符
func CheckSqlValidate(content string) (bool, string) {
	if content == "" {
		return false, ""
	}
	filterString := `exec|execute|insert|select|delete|update|drop|*|chr|mid|master|truncate|
		char|declare|sitename|net user|xp_cmdshell|;|+|create|
		table|from|grant|use|group_concat|column_name|
		information_schema.columns|table_schema|union|where|order|by|count|
		--|//|/|#|or 1 = 1|or '|' or|or'|'or|)or|) or|or(|or (| or|or | and|and |)and|and(|,and|and'`

	arr := strings.Split(filterString, "|")

	for _, s := range arr {
		if index := strings.Index(content, s); index > 0 {
			return true, s
		}
	}
	return false, ""
}

package config

const (
	ContextLogKey        = "cl"
	ERROR_SYS            = "SYS_ERROR"
	ERROR_USER_NOT_EXIST = "PASSPORT_USER_NOT_EXIST" //用户信息不存在
	ERROR_ROWS           = "HAVE_ROWS"               //存在记录
	ERROR_NO_ROWS        = "NO_ROWS"                 //不存在记录
	ERROR_ENERGY_UNABLE  = "ENERGY_UNABLE"           //能量值小于最小值
	ERROR_USER_NO_CARDS  = "USER_NO_CARDS"           //用户无卡片
	ERROR_USE_CARD_LOCK  = "USE_CARD_LOCK"           //频繁使用卡片
	ERROR_NO_SUIT        = "NO_SUIT"                 //无该套装

	DefaultPrivateKeyPwd = "3051634071969152599"
)

type ErrInfo struct {
	Code int
	Msg  string
}

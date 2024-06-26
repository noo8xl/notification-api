package models

// SendTwoStepCodeDto -> via telegram
type SendTwoStepCodeDto struct {
	Code   string `json:"code"`
	ChatID string `json:"chatId"`
}

// TelegramAuthConfig -> auth bot config type
type TelegramAuthConfig struct {
	Token string
	BotID int64
	Key   string
}

// ClietRegistratioDto -> signup a new API client obj
type ClietRegistratioDto struct {
	UserEmail  string `json:"userEmail"`  // cliet email
	DomainName string `json:"domainName"` // name of client domain
}

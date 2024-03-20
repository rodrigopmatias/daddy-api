package models

type Metric struct {
	Id         string `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	TerminalId string `json:"terminalId" gorm:"foreignKey:TerminalRefer;notNull"`
	CreatedAt  int64  `json:"createdAt" gorm:"notNull;default:0"`
}

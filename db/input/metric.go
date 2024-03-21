package input

type Metric struct {
	TerminalId string `json:"terminalId"`
	CreatedAt  int64  `json:"createdAt"`
}

func (c Metric) IsValid() error {
	return nil
}

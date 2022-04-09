package cpu

type CPU struct {
	TotalCores   int          `json:"total_cores"`
	TotalThreads int          `json:"total_threads"`
	Processors   []Processors `json:"processors"`
}

type Processors struct {
	ID           int         `json:"id"`
	TotalCores   int         `json:"total_cores"`
	TotalThreads int         `json:"total_threads"`
	Vendor       string      `json:"vendor"`
	Model        string      `json:"model"`
	Capabilities interface{} `json:"capabilities"`
	Cores        interface{} `json:"cores"`
}

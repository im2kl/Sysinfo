package memory

type Memory struct {
	TotalPhysicalBytes int64       `json:"total_physical_bytes"`
	TotalUsableBytes   int64       `json:"total_usable_bytes"`
	SupportedPageSizes interface{} `json:"supported_page_sizes"`
	Modules            []Modules   `json:"modules"`
}

type Modules struct {
	Label        string `json:"label"`
	Location     string `json:"location"`
	SerialNumber string `json:"serial_number"`
	SizeBytes    int64  `json:"size_bytes"`
	Vendor       string `json:"vendor"`
}

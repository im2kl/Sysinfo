package block

type Block struct {
	TotalSizeBytes int64   `json:"total_size_bytes"`
	Disks          []Disks `json:"disks"`
}

type Disks struct {
	Name                   string `json:"name"`
	SizeBytes              int64  `json:"size_bytes"`
	PhysicalBlockSizeBytes int    `json:"physical_block_size_bytes"`
	DriveType              string `json:"drive_type"`
	Removable              bool   `json:"removable"`
	StorageController      string `json:"storage_controller"`
	BusPath                string `json:"bus_path"`
	Vendor                 string `json:"vendor"`
	Model                  string `json:"model"`
	SerialNumber           string `json:"serial_number"`
	Wwn                    string `json:"wwn"`
	//Partitions             []interface{} `json:"partitions"`
}

package product

type Product struct {
	Family       string `json:"family"`
	Name         string `json:"name"`
	Vendor       string `json:"vendor"`
	SerialNumber string `json:"serial_number"`
	UUID         string `json:"uuid"`
	Sku          string `json:"sku"`
	Version      string `json:"version"`
}

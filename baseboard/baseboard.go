package baseboard

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type Baseboard struct {
	AssetTag     string `json:"asset_tag"`
	SerialNumber string `json:"serial_number"`
	Vendor       string `json:"vendor"`
	Version      string `json:"version"`
	Product      string `json:"product"`
}

func (*Baseboard) Get(b *Baseboard) {

	bb, err := ghw.Baseboard()
	if err != nil {
		fmt.Printf("Error getting net info: %v", err)
	}

	b.AssetTag = bb.AssetTag
	b.SerialNumber = bb.SerialNumber
	b.Vendor = bb.Vendor
	b.Version = bb.Version
	b.Product = bb.Product
}

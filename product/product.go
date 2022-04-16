package product

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type Product struct {
	Family       string `json:"family"`
	Name         string `json:"name"`
	Vendor       string `json:"vendor"`
	SerialNumber string `json:"serial_number"`
	UUID         string `json:"uuid"`
	Sku          string `json:"sku"`
	Version      string `json:"version"`
}

func (*Product) Get(p *Product) {

	pro, err := ghw.Product()
	if err != nil {
		fmt.Printf("Error getting net info: %v", err)
	}

	p.Family = pro.Family
	p.Name = pro.Name
	p.Vendor = pro.Vendor
	p.SerialNumber = pro.SerialNumber
	p.UUID = pro.UUID
	p.Sku = pro.SKU
	p.Version = pro.Version

}

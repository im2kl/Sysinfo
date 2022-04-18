package bios

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type Bios struct {
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
	Date    string `json:"date"`
}

func (*Bios) Get(b *Bios) {
	bi, err := ghw.BIOS()
	if err != nil {
		fmt.Printf("Error getting bios info: %v", err)
	}

	b.Vendor = bi.Vendor
	b.Version = bi.Version
	b.Date = bi.Date

}

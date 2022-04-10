package memory

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type Memory struct {
	TotalPhysicalBytes int64 `json:"total_physical_bytes"`
	TotalUsableBytes   int64 `json:"total_usable_bytes"`
	//SupportedPageSizes interface{} `json:"supported_page_sizes"`
	Modules []Modules `json:"modules"`
}

type Modules struct {
	Label        string `json:"label"`
	Location     string `json:"location"`
	SerialNumber string `json:"serial_number"`
	SizeBytes    int64  `json:"size_bytes"`
	Vendor       string `json:"vendor"`
}

func (x *Memory) Get(m *Memory) {

	mem, err := ghw.Memory()
	if err != nil {
		fmt.Printf("Error getting memory info: %v", err)
	}

	m.TotalPhysicalBytes = mem.TotalPhysicalBytes
	m.TotalUsableBytes = mem.TotalUsableBytes
	//m.SupportedPageSizes = mem.SupportedPageSizes

	md := Modules{}
	for _, y := range mem.Modules {
		//fmt.Printf("%d -- %v\n", x, y.Vendor)
		md.Label = y.Label
		md.Location = y.Location
		md.Vendor = y.Vendor
		md.SerialNumber = y.SerialNumber
		md.SizeBytes = y.SizeBytes

		m.Modules = append(m.Modules, md)
	}
}

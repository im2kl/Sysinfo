package block

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type Block struct {
	TotalSizeBytes int64   `json:"total_size_bytes"`
	Disks          []Disks `json:"disks"`
}

type Disks struct {
	Name      string `json:"name"`
	SizeBytes int64  `json:"size_bytes"`
	//PhysicalBlockSizeBytes int    `json:"physical_block_size_bytes"`
	DriveType         string `json:"drive_type"`
	Removable         bool   `json:"removable"`
	StorageController string `json:"storage_controller"`
	BusPath           string `json:"bus_path"`
	Vendor            string `json:"vendor"`
	Model             string `json:"model"`
	SerialNumber      string `json:"serial_number"`
	Wwn               string `json:"wwn"`
	//Partitions        []interface{} `json:"partitions"`
}

func (*Block) Get(b *Block) {
	blk, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting Chassis info: %v", err)
	}

	b.TotalSizeBytes = int64(blk.TotalPhysicalBytes)

	dk := Disks{}

	for _, y := range blk.Disks {

		dk.Name = y.Name
		dk.SizeBytes = int64(y.SizeBytes)
		//dk.PhysicalBlockSizeBytes = int(y.PhysicalBlockSizeBytes)
		dk.DriveType = y.DriveType.String()
		dk.Removable = y.IsRemovable
		dk.StorageController = y.StorageController.String()
		dk.BusPath = y.BusPath
		dk.Vendor = y.Vendor
		dk.Model = y.Model
		dk.SerialNumber = y.SerialNumber
		dk.Wwn = y.WWN

		b.Disks = append(b.Disks, dk)
	}

}

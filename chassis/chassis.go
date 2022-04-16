package chassis

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type Chassis struct {
	AssetTag        string `json:"asset_tag"`
	SerialNumber    string `json:"serial_number"`
	Type            string `json:"type"`
	TypeDescription string `json:"type_description"`
	Vendor          string `json:"vendor"`
	Version         string `json:"version"`
}

func (*Chassis) Get(c *Chassis) {

	cha, err := ghw.Chassis()
	if err != nil {
		fmt.Printf("Error getting Chassis info: %v", err)
	}

	c.AssetTag = cha.AssetTag
	c.SerialNumber = cha.SerialNumber
	c.Type = cha.Type
	c.TypeDescription = cha.TypeDescription
	c.Vendor = cha.Vendor
	c.Version = cha.Version

}

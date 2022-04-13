package chassis

type Chassis struct {
	AssetTag        string `json:"asset_tag"`
	SerialNumber    string `json:"serial_number"`
	Type            string `json:"type"`
	TypeDescription string `json:"type_description"`
	Vendor          string `json:"vendor"`
	Version         string `json:"version"`
}

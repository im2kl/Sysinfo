package net

type Network struct {
	Nics []Nics `json:"nics"`
}

type Nics struct {
	Name       string `json:"name"`
	MacAddress string `json:"mac_address"`
	IsVirtual  bool   `json:"is_virtual"`
	// Capabilities []interface{} `json:"capabilities"`
}

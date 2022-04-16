package net

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type Network struct {
	Nics []Nics `json:"nics"`
}

type Nics struct {
	Name       string `json:"name"`
	MacAddress string `json:"mac_address"`
	IsVirtual  bool   `json:"is_virtual"`
	// Capabilities []interface{} `json:"capabilities"`
}

func (*Network) Get(n *Network) {

	net, err := ghw.Network()
	if err != nil {
		fmt.Printf("Error getting net info: %v", err)
	}

	nc := Nics{}

	for _, y := range net.NICs {

		nc.Name = y.Name
		nc.MacAddress = y.MacAddress
		nc.IsVirtual = y.IsVirtual

		n.Nics = append(n.Nics, nc)
	}
}

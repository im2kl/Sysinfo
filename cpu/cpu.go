package cpu

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

type CPU struct {
	TotalCores   int          `json:"total_cores"`
	TotalThreads int          `json:"total_threads"`
	Processors   []Processors `json:"processors"`
}

type Processors struct {
	ID           int    `json:"id"`
	TotalCores   int    `json:"total_cores"`
	TotalThreads int    `json:"total_threads"`
	Vendor       string `json:"vendor"`
	Model        string `json:"model"`
	//Capabilities interface{} `json:"capabilities"`
	//Cores        interface{} `json:"cores"`
}

func (*CPU) Get(c *CPU) {

	cpu, err := ghw.CPU()
	if err != nil {
		fmt.Printf("Error getting cpu info: %v", err)
	}

	c.TotalCores = int(cpu.TotalCores)
	c.TotalThreads = int(cpu.TotalThreads)

	ps := Processors{}

	for _, y := range cpu.Processors {

		ps.ID = y.ID
		ps.Model = y.Model
		ps.Vendor = y.Vendor
		ps.TotalThreads = int(y.NumThreads)
		ps.TotalCores = int(y.NumCores)

		c.Processors = append(c.Processors, ps)
	}

}


package gate 

import (
	"github.com/jinzhu/gorm"
)

type GateConfigurationDB struct {
	Engine string `json:"engine"`
	Href string `json:"href"`
	Purge bool `json:"purge"`

	Connection gorm.DB
}

type GateConfiguration struct {
	Database GateConfigurationDB `json:"database"`
}
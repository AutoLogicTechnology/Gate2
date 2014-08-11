
package gate 

import (
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"

    "github.com/jinzhu/gorm"
)

type GateConfigurationQRCode struct {
	WriteToDisk bool `json:"todisk"`
	Path string `json:"path"`
}

type GateConfigurationDB struct {
    Engine string `json:"engine"`
    Href string `json:"href"`
    Setup bool `json:"setup"`

    Connection gorm.DB
}

type GateConfiguration struct {
    Database GateConfigurationDB `json:"database"`
    QRCode GateConfigurationQRCode `json:"qrcodes"`
}

func NewGateConfiguration(cf string) (*GateConfiguration, error) {
	var config *GateConfiguration

	data, err := ioutil.ReadFile(cf)

    if err != nil {
        return nil, errors.New(fmt.Sprintf("Unable to open the given configuration file: %s: %s", cf, err))
    }

    err = json.Unmarshal(data, &config)

    if err != nil {
        return nil, errors.New(fmt.Sprintf("Unable to read configuration format for: %s: %s", cf, err))
    }

    return config, nil 
}
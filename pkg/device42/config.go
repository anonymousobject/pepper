package device42

import (
	"os"

	"github.com/Sirupsen/logrus"

	"github.com/spf13/viper"
)

// Config :
var Config = viper.New()

var (
	// Username :
	Username string
	// Password :
	Password string
	address  string
	// IPRange :
	IPRange string
	// ServiceLevel :
	ServiceLevel string
	// BaseURL :
	BaseURL string
)

// ReadConfig :
func ReadConfig(environment string) error {
	if err := os.MkdirAll("/etc/pepper/provider.d/device42", 0644); err != nil {
		logrus.Warnf("unable to create the /etc/pepper/provider.d/device42/ directory: %v", err)
	}
	Config.SetConfigName(environment)
	Config.SetConfigType("yaml")
	Config.AddConfigPath("/etc/pepper/provider.d/device42")
	if err := Config.ReadInConfig(); err != nil {
		logrus.Fatalf("unable to open device42 config: %v", err)
	}
	// Username :
	Username = Config.GetString("username")
	// Password :
	Password = Config.GetString("password")
	address = Config.GetString("address")
	// IPRange :
	IPRange = Config.GetString("ip_range")
	// ServiceLevel :
	ServiceLevel = Config.GetString("service_level")
	// BaseURL :
	BaseURL = address + "/api/1.0"

	return nil
}

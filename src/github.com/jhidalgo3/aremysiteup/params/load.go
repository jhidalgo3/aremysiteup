package params

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config is a struct that contains config information
type Config struct {
	UrlFile               string
	Sleep, SleepWithError int
	Timeout               int
	OneCheck              bool
	Quiet                 bool
	To                    string
	From                  string
	Urls                  []string
	Mailgun               Mailgun
}

type Mailgun struct {
	Domain, ApiKey, PublicApiKey string
}

func Load(params *Config) {

	viper.AddConfigPath("./configs")
	viper.AddConfigPath("$HOME/configs")

	viper.SetConfigName("aremysiteup")
	viper.SetConfigType("yml")

	//Register Environment Prefix
	viper.SetEnvPrefix("APP")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())

	err := viper.Unmarshal(&params)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}

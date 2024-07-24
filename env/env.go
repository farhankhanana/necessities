package env

import (
	"fmt"
	"log"

	"github.com/gat/necessities/utils"

	"github.com/gat/necessities/validator"
	"github.com/spf13/viper"
)

// InitEnv initiates viper env reader.
func InitEnv(envName, envType string, envPath []string, envModel map[string]interface{}, printEnvValue bool) {
	for _, path := range envPath {
		viper.AddConfigPath(path)
	}

	viper.SetConfigName(envName)
	viper.SetConfigType(envType)

	// replace value from config.env if from OS available
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("[init] read config file failed:", err)
	}

	for name, model := range envModel {
		// marshal to env model
		if err = viper.Unmarshal(model); err != nil {
			log.Fatalln(fmt.Sprintf("[initEnv] viper unmarshal env %s", name), err)
		}
		// validate env requirement model
		if err = validator.StructValidator(model); err != nil {
			log.Fatalln(fmt.Sprintf("[initEnv] missing env %s", name), err)
		}

		// print env value
		if printEnvValue {
			err = utils.PrintStructValue(&model)
		}
	}
}

package config

import (
	"fmt"
	"strings"

	"github.com/go-viper/mapstructure/v2"
	"github.com/spf13/viper"
)

func Load(configFilename string) error {
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	if configFilename != "" {
		v.SetConfigFile(configFilename)
		if err := v.MergeInConfig(); err != nil {
			return fmt.Errorf("failed on config `%s` merging: %w", configFilename, err)
		}
	}

	hooks := []mapstructure.DecodeHookFunc{
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	}
	err := v.Unmarshal(&Values, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
		config.DecodeHook = mapstructure.ComposeDecodeHookFunc(hooks...)
	})
	if err != nil {
		return fmt.Errorf("failed on config `%s` unmarshal: %w", configFilename, err)
	}

	return nil
}

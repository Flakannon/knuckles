package env

import "github.com/spf13/viper"

type EventSourceConfig struct {
	URL  string `mapstructure:"BASE_URL"`
	Name string `mapstructure:"NAME"`
}

func LoadEventSourceConfig() (config EventSourceConfig, err error) {
	if err := viperBindEventSource("KAFKA", &config); err != nil {
		return EventSourceConfig{}, err
	}
	return config, nil
}

func viperBindEventSource(prefix string, config *EventSourceConfig) error {
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()

	viper.BindEnv("BASE_URL")
	viper.BindEnv("NAME")

	return viper.Unmarshal(&config)
}

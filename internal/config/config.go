package config

import (
	"log"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	APP_HOST       string `mapstructure:"APP_HOST"`
	APP_PORT       int    `mapstructure:"APP_PORT"`
	CORS_ORIGINS   string `mapstructure:"CORS_ORIGINS"`
	GEMINI_API_KEY string `mapstructure:"GEMINI_API_KEY"`
}

func InitConfig() (*Config, error) {
	cfg := new(Config)
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigType("env")
	v.SetConfigName(".env")
	v.AutomaticEnv()

	bindEnvStruct(v, cfg)

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("Warning reading config file: %v\n", err)
	}

	err = v.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func bindEnvStruct(v *viper.Viper, s any) {
	val := reflect.ValueOf(s)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return
	}

	typ := val.Type()

	for i := range typ.NumField() {
		field := typ.Field(i)
		tagValue := field.Tag.Get("mapstructure")
		if tagValue != "" {
			v.BindEnv(field.Name, tagValue)
		}
	}
}

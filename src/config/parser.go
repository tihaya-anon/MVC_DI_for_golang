package config

import (
	"MVC_DI/util"
	"log"
	"path"
	"reflect"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// # Parse
//
// Parse Loads a config file into a struct
func Parse[T any](file string, definition *T) error {
	pathStr, fileStr := path.Split(file)
	pathStr = path.Join("..", "resource", pathStr)
	splitFile := strings.Split(fileStr, ".")

	name := splitFile[0]
	ext := splitFile[1]

	viper.SetConfigName(name)
	viper.SetConfigType(ext)
	viper.AddConfigPath(pathStr)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return err
	}

	decoderOption := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeHookFunc(time.RFC3339),
		mapstructure.StringToSliceHookFunc(","),
		func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			if f.Kind() == reflect.Map {
				newMap := make(map[string]interface{})
				for key, value := range data.(map[string]interface{}) {
					newKey := util.SnakeToPascal(key)
					newMap[newKey] = value
				}
				return newMap, nil
			}
			return data, nil
		},
	))

	if err := viper.Unmarshal(definition, decoderOption); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		return err
	}

	return nil
}

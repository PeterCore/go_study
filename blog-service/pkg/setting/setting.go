package setting

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	var configName string = "config"
	path, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	vp := viper.New()
	vp.AddConfigPath(fmt.Sprintf("%s/configs", path))
	vp.SetConfigName(configName)
	vp.SetConfigType("yaml")
	err = vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

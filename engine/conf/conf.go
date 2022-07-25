package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	configName string
	configPath string
	configType string
	configLock sync.Mutex
	configInfo *viper.Viper

	GlobalConf = make(map[string]interface{})
)

func init() {
	fmt.Println("conf init")
	SetConfigFile("//127.0.0.1/suqing/zest/zest.yaml")
	LoadConfig()
}

func LoadConfig() {
	configInfo = viper.New()
	configInfo.AddConfigPath(configPath)
	configInfo.SetConfigName(configName)
	configInfo.SetConfigType(configType)

	if err := configInfo.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Fatal error config file: %s \n", err)
			return
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
}

func SetConfigFile(filePath string) {
	configLock.Lock()
	paths, _ := filepath.Split(filePath)
	fmt.Println(paths)
	if len(filepath.Ext(filePath)) == 0 {
		panic(fmt.Sprintf("SetConfigFile error filePath = %v", filePath))
	}
	types := filepath.Ext(filePath)[1:]
	fmt.Println(types)
	name := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	fmt.Println(name)
	configName = name
	configType = types
	configPath = paths
	configLock.Unlock()
}

// 管理进程配置
func SetGlobalConf(key string, value interface{}) {
	GlobalConf[key] = value
}

func GetGlobalConf(key string) interface{} {
	if _, ok := GlobalConf[key]; !ok {
		return nil
	}
	return GlobalConf[key]
}

func GlobalConfIsSet(key string) bool {
	if _, ok := GlobalConf[key]; ok {
		return true
	}
	return false
}

func Get(key string) interface{} {
	return configInfo.Get(key)
}

func GetBool(key string) bool {
	return configInfo.GetBool(key)
}

func GetFloat64(key string) float64 {
	return configInfo.GetFloat64(key)
}

func GetInt(key string) int {
	return configInfo.GetInt(key)
}

func GetIntSlice(key string) []int {
	return configInfo.GetIntSlice(key)
}

func GetString(key string) string {
	return configInfo.GetString(key)
}

func GetStringMap(key string) map[string]interface{} {
	return configInfo.GetStringMap(key)
}

func GetStringMapString(key string) map[string]string {
	return configInfo.GetStringMapString(key)
}

func GetStringSlice(key string) []string {
	return configInfo.GetStringSlice(key)
}

func GetTime(key string) time.Time {
	return configInfo.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return configInfo.GetDuration(key)
}

func IsSet(key string) bool {
	return configInfo.IsSet(key)
}

func AllSettings() map[string]interface{} {
	return configInfo.AllSettings()
}

package configure

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfigure() {
	viper.SetConfigFile("./farseer.yaml")
	err := viper.ReadInConfig() //找到并读取配置文件
	if err != nil {             // 捕获读取中遇到的error
		log.Printf("[err]读取配置文件farseer.yaml时发生错误: %s \n", err)
	}
}

// GetString 获取配置
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt 获取配置
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetInt64 获取配置
func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

// GetBool 获取配置
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetSubNodes 获取所有子节点
func GetSubNodes(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// SetDefault 设置配置的默认值
func SetDefault(key string, value interface{}) {
	viper.SetDefault(key, value)
}

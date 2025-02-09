package configure

import (
	"strings"
)

var configurationBuilder = newConfigurationBuilder()

type config struct {
	def            map[string]any    // 默认配置
	envKeyReplacer *strings.Replacer // 环境变量替换
	configProvider []IConfigProvider // 配置提供者
}

func newConfigurationBuilder() config {
	return config{
		def:            make(map[string]any),
		configProvider: []IConfigProvider{},
	}
}

// AddYamlFile 设置yaml文件配置
func (c *config) AddYamlFile(configFile string) {
	var yConfig IConfigProvider = NewYamlConfig(configFile)
	c.configProvider = append(c.configProvider, yConfig)
}

// AddEnvironmentVariables 加载环境变量
func (c *config) AddEnvironmentVariables() {
	var envConfig IConfigProvider = NewEnvConfig()
	c.configProvider = append([]IConfigProvider{envConfig}, c.configProvider...)
}

// SetEnvKeyReplacer 环境变量替换
func (c *config) SetEnvKeyReplacer(r *strings.Replacer) {
	c.envKeyReplacer = r
}

// Build 找到并读取配置文件
func (c *config) Build() error {
	for _, provider := range c.configProvider {
		err := provider.LoadConfigure()
		if err != nil {
			return err
		}
	}
	return nil
}

// Get 获取配置
func (c *config) Get(key string) any {
	// 遍历配置提供者
	for _, provider := range c.configProvider {
		v, exists := provider.Get(key)
		if exists {
			return v
		}
	}

	// 是否有默认配置
	val, exists := c.def[key]
	if exists {
		return val
	}

	return nil
}

// GetSubNodes 获取所有子节点
func (c *config) GetSubNodes(key string) map[string]any {
	// 这里需要倒序获取列表，利用后面覆盖前面的方式来获取
	m := make(map[string]any)
	// 先添加默认值
	prefixKey := key + "."
	for k, v := range c.def {
		if strings.HasPrefix(k, prefixKey) {
			m[k[len(prefixKey):]] = v
		}
	}

	// 再添加yaml、环境变量
	for i := len(c.configProvider) - 1; i >= 0; i-- {
		subMap, exists := c.configProvider[i].GetSubNodes(key)
		if exists {
			for k, v := range subMap {
				m[k] = v
			}
		}
	}

	return m
}

// GetSlice 获取数组
func (c *config) GetSlice(key string) []string {
	// 遍历配置提供者
	for _, provider := range c.configProvider {
		v, exists := provider.Get(key)
		if exists {
			m, isOk := v.([]any)
			if isOk {
				var arr []string
				for _, s := range m {
					arr = append(arr, s.(string))
				}
				return arr
			}
		}
	}
	return []string{}
}

// GetSliceNodes 获取数组节点
func (c *config) GetSliceNodes(key string) []map[string]any {
	// 遍历配置提供者
	for _, provider := range c.configProvider {
		v, exists := provider.Get(key)
		if exists {
			arr, isOk := v.([]any)
			if isOk {
				var sliceNodes []map[string]any
				for _, node := range arr {
					sliceNodes = append(sliceNodes, node.(map[string]any))
				}
				return sliceNodes
			}
		}
	}
	return []map[string]any{}
}

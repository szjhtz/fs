package configure

type IConfigProvider interface {
	// LoadConfigure 加载配置
	LoadConfigure() error
	// Get 读取配置
	Get(key string) (any, bool)
	// GetSubNodes 获取所有子节点
	GetSubNodes(key string) (map[string]any, bool)
}

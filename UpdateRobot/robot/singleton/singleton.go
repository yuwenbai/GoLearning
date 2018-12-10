package singleton

var _self *Singleton

//Singleton 类型定义
type Singleton struct {
	maintenance bool
}

//Instance 实例
func Instance() *Singleton {
	if _self == nil {
		_self = new(Singleton)
		return _self
	}
	return _self
}

//SetMaintenanceStatus 全局维护状态
func (o *Singleton) SetMaintenanceStatus(s bool) {
	_self.maintenance = s
}

//GetMaintenanceStatus 获取维护状态
func (o *Singleton) GetMaintenanceStatus() bool {
	return _self.maintenance
}

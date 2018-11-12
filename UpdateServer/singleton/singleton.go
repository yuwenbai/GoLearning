package singleton

var _self *Singleton

type Singleton struct {
	maintenance bool
}

func Instance() *Singleton {
	if _self == nil {
		_self = new(Singleton)
		return _self
	}
	return _self
}

func (o *Singleton) SetMaintenanceStatus(s bool) {
	_self.maintenance = s
}

func (o *Singleton) GetMaintenanceStatus() bool {
	return _self.maintenance
}

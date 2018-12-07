package main

import "time"

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo

//UpdateInfo 信息实例
type UpdateInfo struct {
	VersionID   string `json:"VersionID"`
	VersionName string `json:"VersionName"`
	VersionInfo string `json:"VersionInfo"`
	FileSize    int64  `json:"FileSize"`
	FileName    string `json:"FileName"`
	NeedInstall bool   `json:"NeedInstall"`
	NeedUpdate  bool   `json:"NeedUpdate"`
}

//UpdateInfos 信息数组
type UpdateInfos []UpdateInfo

//Result http
type Result struct {
	Ret    int
	Reason string
	Data   interface{}
}

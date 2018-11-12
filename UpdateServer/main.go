package main

import (
	"net/http"

	"github.com/golang/glog"
)

func main() {

	glog.Info("updateserver start！")
	router := NewRouter()

	glog.Fatal(http.ListenAndServe(":8080", router))

	// // glog.Errorf("Error Line:%d\n", i+1)
	// glog.Errorf("updateserver start！")
	// glog.Infof("updateserver start！")
	// glog.Warningf("updateserver start！")
	// // glog.Warningf("Warning Line:%d\n", i+1)
	// glog.Flush()
}

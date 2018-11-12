package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/golang/glog"
)

//DoDiffApk 差异校对
func DoDiffApk() error {
	pwd, _ := os.Getwd()
	src := pwd + string(os.PathSeparator) + constantPathTempapk
	err := filepath.Walk(src, func(fileNamePath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		println(fileNamePath)
		apknamenosuffix := getFileNameNoSuffix(fileNamePath)
		spkname, version := splitApkName(apknamenosuffix)
		DiffApk(fileNamePath, spkname, version)
		return nil
	})
	if err != nil {
		// fmt.Printf("DoDiffApk.Walk() returned %v\n", err)
		return err
	}
	return nil
}

//DiffApk 差异对比
func DiffApk(src, apkName, versionName string) error {
	pwd, _ := os.Getwd()
	srcPatchPath := pwd + string(os.PathSeparator) + constantPathPatch + apkName + "/"
	srcVersionPath := pwd + string(os.PathSeparator) + constantPathVersions + apkName + "/"

	err := os.Mkdir(srcPatchPath, 0666)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir(srcVersionPath, 0666)
	if err != nil {
		fmt.Println(err)
	}

	lock := &sync.Mutex{}
	var wg sync.WaitGroup
	err = filepath.Walk(srcVersionPath, func(filename string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fmt.Println("FILE: " + filename)
		filenameOnly := getFileNameNoSuffix(filename)
		newPathName := srcPatchPath + filenameOnly + ".patch"

		// var execmd = exec.Command("bsdiff-win.exe", filename, src, newPathName)
		// exeerr := execmd.Run()
		// if exeerr != nil {
		// 	log.Fatal(exeerr)
		// }

		wg.Add(1)
		go func(oldFileName, newFileName, generateFileWithPath string) {
			defer wg.Done()
			lock.Lock()
			var execmd = exec.Command("bsdiff-win.exe", oldFileName, newFileName, generateFileWithPath)
			exeerr := execmd.Run()
			if exeerr != nil {
				glog.Fatal(exeerr)
			}
			lock.Unlock()
		}(filename, src, newPathName)

		return nil
	})

	wg.Wait()

	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	err = os.Rename(src, srcVersionPath+versionName+".apk")
	return nil
}

func getFileNameNoSuffix(filewithsuffix string) string {
	fileNameNoPath := filepath.Base(filewithsuffix)
	fileSuffix := path.Ext(fileNameNoPath)
	filenameOnly := strings.TrimSuffix(fileNameNoPath, fileSuffix)
	return filenameOnly
}

//@eg apknamenosuffix   aaaa_111 ------> aaaa and 111
func splitApkName(apknamenosuffix string) (string, string) {
	s := strings.Split(apknamenosuffix, "_")
	return s[0], s[1]
}

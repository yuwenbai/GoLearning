package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"./singleton"
	"./utillog"
)

//DoDiffApk 差异校对
func DoDiffApk() (string, error) {

	var highestVersion = ""
	pwd, _ := os.Getwd()
	src := pwd + string(os.PathSeparator) + constantPathTempapk
	err := filepath.Walk(src, func(fileNamePath string, f os.FileInfo, err error) error {
		if strings.Contains(fileNamePath, constantMacFileStore) == true {
			return nil
		}
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fmt.Printf("fdfdf %s fileNamePath ", fileNamePath)
		apknamenosuffix := getFileNameNoSuffix(fileNamePath)
		fmt.Printf("fdfdf %s ", apknamenosuffix)
		spkname, version := splitApkName(apknamenosuffix)
		DiffApk(fileNamePath, spkname, version)
		highestVersion = version
		return nil
	})
	if err != nil {
		// fmt.Printf("DoDiffApk.Walk() returned %v\n", err)
		return highestVersion, err
	}
	return highestVersion, nil
}

//DiffApk 差异对比
func DiffApk(src, apkName, versionID string) error {
	pwd, _ := os.Getwd()
	srcPatchPath := pwd + string(os.PathSeparator) + constantPathPatch + apkName + "/"
	srcVersionPath := pwd + string(os.PathSeparator) + constantPathVersions + apkName + "/"

	err := os.Mkdir(srcPatchPath, 0777)
	if err != nil {
		fmt.Println(err)
	}

	err = os.Mkdir(srcVersionPath, 0777)
	if err != nil {
		fmt.Println(err)
	}

	if singleton.Instance().AppIsIos(apkName) {
		appName := singleton.Instance().GetPackageFullName(apkName, versionID)
		iosSrcVersionPath := pwd + string(os.PathSeparator) + constantPathVersions
		err = os.Rename(src, iosSrcVersionPath+appName)
		if err != nil {
			fmt.Printf("rename failed %v\n", err)
		}
	} else {
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

			wg.Add(1)
			go func(oldFileName, newFileName, generateFileWithPath string) {
				defer wg.Done()
				lock.Lock()
				utillog.Instance().Info("oldFileName " + oldFileName + " newFileName " + newFileName + " generateFileWithPath " + generateFileWithPath)
				var execmd = exec.Command("bsdiff", oldFileName, newFileName, generateFileWithPath)

				var out bytes.Buffer
				var stderr bytes.Buffer
				execmd.Stdout = &out
				execmd.Stderr = &stderr
				err := execmd.Run()
				if err != nil {
					fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
					return
				}
				fmt.Println("Result: " + out.String())

				// output, err := execmd.CombinedOutput()
				// if err != nil {
				// 	fmt.Println(fmt.Sprint(err) + ": " + string(output))
				// 	utillog.Instance().Info(output)
				// 	utillog.Instance().Info(err)
				// 	utillog.Instance().Fatal(err)
				// 	return
				// }
				// fmt.Println(string(output))

				lock.Unlock()
			}(filename, src, newPathName)

			return nil
		})

		wg.Wait()

		if err != nil {
			fmt.Printf("filepath.Walk() returned %v\n", err)
		}
		err = os.Rename(src, srcVersionPath+versionID+".apk")
		if err != nil {
			fmt.Printf("rename failed %v\n", err)
		}
	}
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
	fmt.Println(apknamenosuffix)
	s := strings.Split(apknamenosuffix, "_")
	fmt.Println(s)
	return s[0], s[1]
}

//GetCurrentDirectory 获取当前工作目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

//DoModifyPlist modify plist
func DoModifyPlist(version string) error {

	d := GetCurrentDirectory()

	fmt.Println(d)
	tFile := d + string(os.PathSeparator) + "workspace/versions/script.py"

	verURL := version + ".ipa"

	fmt.Println(tFile)
	cmd := exec.Command("python", tFile, verURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	log.Println(err)
	return err
}

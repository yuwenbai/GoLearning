package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"./singleton"
)

//Index 启动页 顺便请求最新版本号信息
func Index(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte(tpl))

	t, err := template.ParseFiles("template/html/login/index.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

//TodoIndex 测试
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

//GetAllPackageNames 获取所有表名
func GetAllPackageNames(w http.ResponseWriter, r *http.Request) {
	nameArray := LookUpGroupBy("ApkName")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err := json.NewEncoder(w).Encode(nameArray); err != nil {
		panic(err)
	}
}

//GetCurrentVersion 获取版本
func GetCurrentVersion(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	r.ParseForm()
	apkName := r.Form["apkName"]
	var array, ret = UpdateRecord(apkName[0])
	if len(array) > 0 && true == ret {
		info := UpdateInfo{NeedInstall: true, VersionId: array[0].Version_id, VersionName: array[0].Version_name, VersionInfo: array[0].Version_info, FileSize: 0, FileName: ""}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(info); err != nil {
			panic(err)
		}
	}
}

//TodoCreate 解析post
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	r.ParseForm()
	todoID := r.Form["todoId"]
	fmt.Println("Form: ", r.Form)
	fmt.Println("Path: ", r.URL.Path)
	fmt.Println(todoID)

	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, constantIoMax))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := r.Body.Close(); err != nil {
	// 	panic(err)
	// }
	// if err := json.Unmarshal(body, &todo); err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(422) // unprocessable entity
	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
	// 		panic(err)
	// 	}
	// }

	// // t := RepoCreateTodo(todo)
	// t := todo
	// w.Header().Set("ontent-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(t); err != nil {
	// 	panic(err)
	// }
}

//GeneratePatch 生成差异
func GeneratePatch(w http.ResponseWriter, r *http.Request) {
	err := DoDiffApk()
	if err != nil {
		OutputJSON(w, 0, err.Error(), nil)
		return
	}
	OutputJSON(w, 0, "生成成功，维护状态已取消", nil)
	singleton.Instance().SetMaintenanceStatus(false)
}

//receiver apk
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm

		fmt.Println(r.FormValue("apk_name"))
		fmt.Println(r.FormValue("version_num"))
		fmt.Println(r.FormValue("version_name"))
		fmt.Println(r.FormValue("version_content"))

		apkNameString := r.FormValue("apk_name")
		versionID := r.FormValue("version_num")

		if apkNameString == "" || versionID == "" {
			OutputJSON(w, 0, "检查包名或者版本号", nil)
			return
		}

		//这里有问题 要改成order by 查找最大版本号与当前版本号进行大小比对
		var array, _ = UpdateRecord(apkNameString)

		if len(array) > 0 && strings.Compare(array[0].Version_id, versionID) >= 0 {
			OutputJSON(w, 0, "版本号需高于当前版本号，请检查版本号", nil)
			return
		}
		pwd, _ := os.Getwd()
		path := pwd + string(os.PathSeparator) + constantPathTempapk

		fmt.Println("insert Record path " + path)
		//get the *fileheaders
		files := m.File["uploadfile"]
		fmt.Println(files)
		for i := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			//@eg: apkNameString + "_" + versionID "xxxxname_111"

			var filepath = path + apkNameString + "_" + versionID
			dst, err := os.Create(filepath)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		singleton.Instance().SetMaintenanceStatus(true)

		fmt.Println("insert Record")
		//success read file complete refresha the new versioninfo to db
		InsertRecord(apkNameString, versionID, r.FormValue("version_name"), r.FormValue("version_content"))

		OutputJSON(w, 0, "上传完成，请返回上一级界面生成差异", nil)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//download apk
func handerGetFile(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// todoId := vars["todoId"]
	vars := r.URL.Query()
	apkName, ok := vars["name"]
	if !ok {
		fmt.Printf("param a does not exist\n")
	} else {
		fmt.Printf("param a value is [%s]\n", apkName[0])
	}
	// r.ParseForm() //解析参数，默认是不会解析的

	var array, ret = UpdateRecord(apkName[0])
	if len(array) > 0 && ret == true {
		pwd, _ := os.Getwd()
		// des := pwd + string(os.PathSeparator) + r.URL.Path[1:len(r.URL.Path)]
		des := pwd + string(os.PathSeparator) + constantPathVersions + array[0].Apk_name + "/" + array[0].Version_id + ".apk"
		desStat, err := os.Stat(des)
		if err != nil {
			http.NotFoundHandler().ServeHTTP(w, r)
		} else if desStat.IsDir() {
			http.NotFoundHandler().ServeHTTP(w, r)
		} else {
			fileData, err := ioutil.ReadFile(des)
			if err != nil {
				log.Println("Read File Err:", err.Error())
			} else {
				w.Header().Set("Content-Type", "application/zip")
				w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", array[0].Apk_name+".apk"))
				w.Write(fileData)
			}
		}
	}

}

//CheckUpdateInfoJSON json接口 检查更新
func CheckUpdateInfoJSON(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	apkName, ok := vars["name"]
	if !ok {
		fmt.Printf("param a does not exist\n")
	} else {
		fmt.Printf("param a value is [%s]\n", apkName[0])
	}
	version, ok := vars["version"]
	if !ok {
		fmt.Printf("param a does not exist\n")
	} else {
		fmt.Printf("param a value is [%s]\n", version[0])
	}
	if singleton.Instance().GetMaintenanceStatus() == true {
		emp1 := UpdateInfo{NeedUpdate: false, NeedInstall: false, VersionId: "", VersionName: "", VersionInfo: "", FileSize: 0, FileName: ""}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(emp1); err != nil {
			panic(err)
		}
	} else {
		if len(version) > 0 && len(apkName) > 0 {
			var array, ret = UpdateRecord(apkName[0])
			if ret == true && len(array) > 0 {
				versionid := array[0].Version_id
				if versionid == version[0] {
					emp1 := UpdateInfo{NeedUpdate: false, NeedInstall: false, VersionId: versionid, VersionName: "", VersionInfo: "", FileSize: 0, FileName: ""}
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusOK)
					if err := json.NewEncoder(w).Encode(emp1); err != nil {
						panic(err)
					}
				} else {
					versionname := array[0].Version_name
					versioninfo := array[0].Version_info
					//查找文件获取文件信息
					pwd, _ := os.Getwd()
					patch := pwd + string(os.PathSeparator) + constantPathPatch
					newPathName := patch + apkName[0] + string(os.PathSeparator) + version[0] + ".patch"
					fileInfo, err := os.Stat(newPathName)
					if err != nil { //强制更新 完整包
						localfile := pwd + string(os.PathSeparator) + constantPathVersions + apkName[0] + string(os.PathSeparator) + versionid + ".apk"
						fileInfo, err := os.Stat(localfile)
						if err != nil {

						} else {
							fileName := constantPathFileserverVersions + apkName[0] + "/" + versionid + ".apk"
							emp1 := UpdateInfo{NeedUpdate: true, NeedInstall: true, VersionId: versionid, VersionName: versionname, VersionInfo: versioninfo, FileSize: fileInfo.Size(), FileName: fileName}
							w.Header().Set("Content-Type", "application/json; charset=UTF-8")
							w.WriteHeader(http.StatusOK)
							if err := json.NewEncoder(w).Encode(emp1); err != nil {
								panic(err)
							}
						}
					} else {
						var size = fileInfo.Size()
						fileName := constantPathFileserverPatch + apkName[0] + "/" + version[0] + ".patch"
						// UpdateInfos := UpdateInfos{
						// 	UpdateInfo{Needinstall: false, VersionId: versionid, VersionName: versionname, VersionInfo: versioninfo, FileSize: size, FileName: fileName},
						// }
						emp1 := UpdateInfo{NeedUpdate: true, NeedInstall: false, VersionId: versionid, VersionName: versionname, VersionInfo: versioninfo, FileSize: size, FileName: fileName}
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusOK)
						if err := json.NewEncoder(w).Encode(emp1); err != nil {
							panic(err)
						}
					}
				}
			}

		}
	}

}

//OutputJSON 返回json测试
func OutputJSON(w http.ResponseWriter, ret int, reason string, i interface{}) {
	out := &Result{ret, reason, i}
	b, err := json.Marshal(out)
	if err != nil {
		return
	}
	w.Write(b)
}

// const tpl = `<html><head><title>上传文件</title></head><body><form enctype="multipart/form-data" action="/upload" method="post"> <input type="file" name="uploadfile" /> <input type="hidden" name="token" value="{...{.}...}"/> <input type="submit" value="upload" /></form></body></html>`

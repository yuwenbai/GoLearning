package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"./utillog"
	_ "github.com/go-sql-driver/mysql"
)

//UpdateVersion 版本更新相关信息
type UpdateVersion struct {
	APPName     string
	VersionID   string
	VersionName string
	VersionInfo string
	VersionTime string
}

//InsertRecord 插入记录
func InsertRecord(ApkName, VersionID, VersionName, VersionInfo, PackageType string) bool {
	// username: root; password: 123456; database: test
	// db, err := sql.Open("mysql", "root:111111@/updatateserverdb")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	var connString, err = GenerateConnectSQLDB()
	if err != nil {
		utillog.Instance().Fatal("GenerateConnectSQLDB failed:", err.Error())
	}

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建

	cmd := fmt.Sprintf("INSERT INTO %s (ApkName, VersionID, VersionName, VersionInfo, CreatedOn) VALUES ('%s',  '%s', '%s', '%s', '%s')", "UpdateServerDB_Version", ApkName, VersionID, VersionName, VersionInfo, time.Now().Format("2006-01-02 15:04:05"))
	ret, err := db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret.RowsAffected()) //更新的条目数
	return true
}

//LookUpGroupBy 获取所有表名
func LookUpGroupBy(columnName string) []string {
	// db, err := sql.Open("mysql", "root:111111@/updatateserverdb")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	var connString, err = GenerateConnectSQLDB()
	if err != nil {
		utillog.Instance().Fatal("GenerateConnectSQLDB failed:", err.Error())
	}

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cmd := fmt.Sprintf("select distinct %s from %s", columnName, "UpdateServerDB_Version")
	rows, err := db.Query(cmd)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// cloumns, err := rows.Columns()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var strdata = make([]string, 0)

	cols, err := rows.Columns()
	if err != nil {
		utillog.Instance().Fatal(err.Error())
		return nil
	}
	vals := make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		vals[i] = new(interface{})
		if i != 0 {
			fmt.Print("\t")
		}
		fmt.Print(cols[i])
	}
	fmt.Println()

	var tempvalue string

	cols, err = rows.Columns()
	if err != nil {
		utillog.Instance().Fatal(err.Error())
	}

	length := len(cols)

	fmt.Println(length)
	for rows.Next() {

		err = rows.Scan(&tempvalue)
		if err != nil {
			utillog.Instance().Fatal(err.Error())
		}
		strdata = append(strdata, tempvalue)
	}
	return strdata

}

//GenerateConnectSQLDB 连接数据库
func GenerateConnectSQLDB() (string, error) {
	//导入配置文件
	configMap := InitConfig(constantPathConfig + "db_configuration.txt")

	caddr := configMap["host"]
	cport := configMap["port"]
	cuser := configMap["user"]
	cpassword := configMap["password"]
	cdbname := configMap["dbname"]

	var server = caddr
	var port = cport
	var user = cuser
	var password = cpassword
	var database = cdbname

	//连接字符串
	// "root:123456@tcp(192.168.2.225:3306)/dbname?charset=utf8"
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, server, port, database)
	// connString := "root:111111@tcp(127.0.0.1:3306)/updateserverdb?charset=utf8"
	return connString, nil

}

//UpdateRecord 更新数据库 获取对应包名所有信息 按versionId排序
func UpdateRecord(apkName string) ([]*UpdateVersion, bool) {

	var connString, err = GenerateConnectSQLDB()

	fmt.Println(connString)
	if err != nil {
		utillog.Instance().Fatal("GenerateConnectSQLDB failed:", err.Error())
	}

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// cmd := fmt.Sprintf("SELECT * FROM %s where ApkName='%s' ORDER BY right ('0000000000'+VersionID,10) desc", "UpdateServerDB_Version", apkName)
	cmd := fmt.Sprintf("SELECT * FROM %s where ApkName='%s' ORDER BY VersionID+0 desc", "UpdateServerDB_Version", apkName)
	var sqlArray, err1 = exeUpdateSQL(db, cmd)

	for index, item := range sqlArray {
		fmt.Println(index)
		fmt.Println(item)
	}
	if err1 != nil {
		utillog.Instance().Fatal("failed:", err1.Error())
		return nil, false
	}
	return sqlArray, true
}

func exeUpdateSQL(db *sql.DB, cmd string) ([]*UpdateVersion, error) {
	rows, err := db.Query(cmd)
	if err != nil {
		utillog.Instance().Fatal("failed:", err.Error())
		return nil, err
	}
	defer rows.Close()

	var retRowData []*UpdateVersion
	//遍历每一行
	for rows.Next() {
		var row = new(UpdateVersion)
		rows.Scan(&row.APPName, &row.VersionID, &row.VersionName, &row.VersionInfo, &row.VersionTime)
		retRowData = append(retRowData, row)
	}

	if rows.Err() != nil {
		return retRowData, rows.Err()
	}
	return retRowData, nil
}
func printValue(pval *interface{}) {
	switch v := (*pval).(type) {
	case nil:
		fmt.Print("NULL")
	case bool:
		if v {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	case []byte:
		fmt.Print(string(v))
	case time.Time:
		fmt.Print(v.Format("2006-01-02 15:04:05.999"))
	default:
		fmt.Print(v)
	}
}

package main

import (
	"database/sql"
	"fmt"
	"time"

	"./utillog"
	_ "github.com/denisenkom/go-mssqldb"
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

	var connString, err = GenerateConnectSQLDB()
	if err != nil {
		utillog.Instance().Fatal("GenerateConnectSQLDB failed:", err.Error())
	}
	//建立连接
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		utillog.Instance().Fatal("Open Connection failed:", err.Error())
	}
	defer conn.Close()

	cmd := fmt.Sprintf("INSERT INTO %s (ApkName, VersionID, VersionName, VersionInfo, CreatedOn) VALUES ('%s',  '%s', '%s', '%s', '%s')", "dbo.UpdateServerDB_Version", ApkName, VersionID, VersionName, VersionInfo, time.Now().Format("2006-01-02 15:04:05"))
	err = exeSQL(conn, cmd)
	if err != nil {
		utillog.Instance().Fatal(err.Error())
		return false
	}
	return true
}

//LookUpGroupBy 获取所有表名
func LookUpGroupBy(columnName string) []string {

	var connString, err = GenerateConnectSQLDB()
	if err != nil {
		utillog.Instance().Fatal("GenerateConnectSQLDB failed:", err.Error())
	}
	//建立连接
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		utillog.Instance().Fatal("Open Connection failed:", err.Error())
	}
	defer conn.Close()
	// cmd := fmt.Sprintf("select '%s' from %s group by %s", columnName, "dbo.UpdateServerDB_Version", columnName)
	// cmd := fmt.Sprintf("select %s from %s group by %s", columnName, "dbo.UpdateServerDB_Version", columnName)
	cmd := fmt.Sprintf("select distinct %s from %s", columnName, "dbo.UpdateServerDB_Version")
	rows, err := conn.Query(cmd)
	if err != nil {
		utillog.Instance().Fatal(err.Error())
	}
	//遍历每一行

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
	connString := fmt.Sprintf("server=%s;port=%s;database=%s;user id=%s;password=%s", server, port, database, user, password)
	return connString, nil
}

//UpdateRecord 更新数据库 获取对应包名所有信息 按versionId排序
func UpdateRecord(apkName string) ([]*UpdateVersion, bool) {

	var connString, err = GenerateConnectSQLDB()
	if err != nil {
		utillog.Instance().Fatal("GenerateConnectSQLDB failed:", err.Error())
	}
	//建立连接
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		utillog.Instance().Fatal("Open Connection failed:", err.Error())
	}
	defer conn.Close()

	cmd := fmt.Sprintf("SELECT * FROM %s where ApkName='%s' ORDER BY right ('0000000000'+VersionID,10) desc", "dbo.UpdateServerDB_Version", apkName)
	var sqlArray, err1 = exeUpdateSQL(conn, cmd)
	if err1 != nil {
		utillog.Instance().Fatal("failed:", err1.Error())
		return nil, false
	}
	return sqlArray, true
}

// //LookUpAllDBInfo 获取所有数据
// func LookUpAllDBInfo(apkName string) ([]*UpdateVersion, bool) {

// 	var connString, err = GenerateConnectSQLDB()
// 	if err != nil {
// 		utillog.Instance().Fatal("GenerateConnectSQLDB failed:", err.Error())
// 	}
// 	//建立连接
// 	conn, err := sql.Open("mssql", connString)
// 	if err != nil {
// 		utillog.Instance().Fatal("Open Connection failed:", err.Error())
// 	}
// 	defer conn.Close()

// 	cmd := fmt.Sprintf("SELECT * FROM %s where ApkName='%s' ORDER BY CreatedOn desc", "dbo.UpdateServerDB_Version", apkName)
// 	var sqlArray, err1 = exeUpdateSQL(conn, cmd)
// 	if err1 != nil {
// 		utillog.Instance().Fatal("failed:", err1.Error())
// 		return nil, false
// 	}
// 	fmt.Println(sqlArray)
// 	return sqlArray, true
// }

//exeSQL 执行sql语句
func exeSQL(db *sql.DB, cmd string) error {
	rows, err := db.Query(cmd)
	if err != nil {
		utillog.Instance().Fatal("failed:", err.Error())
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		utillog.Instance().Fatal("failed:", err.Error())
		return err
	}
	if cols == nil {
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
	for rows.Next() {
		err = rows.Scan(vals...)
		if err != nil {
			utillog.Instance().Fatal("failed:", err.Error())
			continue
		}
		for i := 0; i < len(vals); i++ {
			if i != 0 {
				fmt.Print("\t")
			}
			printValue(vals[i].(*interface{}))
		}
		fmt.Println()

	}
	if rows.Err() != nil {
		return rows.Err()
	}
	return nil
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

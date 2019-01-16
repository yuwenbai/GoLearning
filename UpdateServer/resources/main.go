package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./singleton"
	"./utillog"
	"github.com/garyburd/redigo/redis"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "./customer"
	testins "./testins"
)

type sparseBundleHeader struct {
	CFBundleName        string `plist:"kind"`
	CFBundleExecuteable string `plist:"url"`
}

type array []interface{}

// Person test
type Person struct {
	//对应id表字段
	ApkName string `db:"ApkName"`
	//对应name表字段
	VersionID string `db:"VersionID"`
	//对应age表字段
	VersionName string `db:"VersionName"`
	//对应rmb表字段
	VersionInfo string `db:"VersionInfo"`
	//对应rmb表字段
	CreateOn string `db:"CreateOn"`
}

// GetAll 获取
func GetAll() {
	//先看看redis里有没有数据
	conn, err := redis.Dial("tcp", "66.112.212.102:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	reply, err := conn.Do("lrange", "mlist", 0, -1)
	pkeys, _ := redis.Strings(reply, err)
	fmt.Println(pkeys)

	if len(pkeys) > 0 {
		//如果有
		fmt.Println("从redis获得数据")

		// 从redis里直接读取
		for _, key := range pkeys {
			retStrs, _ := redis.Strings(conn.Do("hgetall", key))
			//fmt.Println(retStrs)
			fmt.Printf("{%s %s %s}\n", retStrs[1], retStrs[3], retStrs[5])
		}

	} else {
		//如果没有
		fmt.Println("从mysql获得数据")

		//查询数据库
		var connString, err = GenerateConnectSQLDB()
		if err != nil {
			utillog.Instance().Fatal("GenerateConnectSQLDB failed:", err.Error())
		}
		//建立连接
		db, err := sqlx.Open("mssql", connString)
		if err != nil {
			utillog.Instance().Fatal("Open Connection failed:", err.Error())
		}
		defer db.Close()

		var persons []Person
		db.Select(&persons, "select ApkName,VersionID,VersionName,VersionInfo from dbo.UpdateServerDB_Version")
		fmt.Println(persons)

		//写入redis并且设置过期时间
		for _, p := range persons {
			//将p以hash形式写入redis
			_, e1 := conn.Do("hmset", p.ApkName, "ApkName", p.ApkName, "VersionID", p.VersionID, "VersionName", p.VersionName, "VersionInfo", p.VersionInfo)

			//将这个hash的key加入mlist
			_, e2 := conn.Do("rpush", "mlist", p.ApkName)

			//设置过期时间
			_, e3 := conn.Do("expire", p.ApkName, 60)
			_, e4 := conn.Do("expire", "mlist", 60)

			if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
				fmt.Println(p.VersionID, "写入失败", e1, e2, e3, e4)
			} else {
				fmt.Println(p.VersionID, "写入成功")
			}
		}
	}
}

func main() {

	singleton.Instance().SetMaintenanceStatus(false)
	testins.TestX()
	// s := Base64String("test")
	// f, err := os.Open("test.plist")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()
	// dco := plist.NewDecoder(f)

	// var bval map[string]interface{}

	// dco.Decode(&bval)
	// fmt.Println(bval)

	// var aval []interface{}
	// dco.Decode(&aval)
	// fmt.Println(aval)

	// fmt.Println("ffffffffffffffffffffff begin")
	// for k, v := range bval {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		fmt.Println(k, "is string", vv)
	// 	case float64:
	// 		fmt.Println(k, "is float64", vv)
	// 	case []interface{}:
	// 		fmt.Println(k, "is an array:")
	// 		for i, u := range vv {
	// 			fmt.Println(i, u)
	// 		}
	// 	default:
	// 		fmt.Println(k, "is of a type I don't know how to handle")
	// 	}
	// }

	// fmt.Println("ffffffffffffffffffffff")

	// fw, err := os.Open("testnew.plist")
	// defer fw.Close()
	// enco := plist.NewEncoder(f)
	// enco.Encode(bval)

	// dco.Unmarshal(data)

	// data, err := plist.Marshal(&s, plist.OpenStepFormat)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Property List:", string(data))

	// var decoded Base64String
	// _, err = plist.Unmarshal(dco, &decoded)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Raw Data:", string(decoded))

	// c, err := redis.Dial("tcp", "127.0.0.1:6379")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //密码授权
	// c.Do("AUTH", "123456")
	// c.Do("SET", "a", 134)
	// a, err := redis.Int(c.Do("GET", "a"))

	// fmt.Println(a)

	// defer c.Close()

	// //创建输出日志文件
	// logFile, err := os.Create("./log/" + time.Now().Format("20060102") + ".txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //创建一个Logger
	// //参数1：日志写入目的地
	// //参数2：每条日志的前缀
	// //参数3：日志属性
	// loger := utillog.Instance().New(logFile, "test_", utillog.Instance().Lmicroseconds|utillog.Instance().Llongfile)

	// //Flags返回Logger的输出选项
	// fmt.Println(loger.Flags())

	// //SetFlags设置输出选项
	// // loger.SetFlags(utillog.Instance().Ldate | utillog.Instance().Ltime | utillog.Instance().Lshortfile)

	// //返回输出前缀
	// fmt.Println(loger.Prefix())

	// //设置输出前缀
	// loger.SetPrefix("test_")

	// //输出一条日志
	// loger.Output(2, "打印一条日志信息")

	// //格式化输出日志
	// loger.Printf("第%d行 内容:%s", 11, "我是错误")

	// //等价于print();os.Exit(1);
	// loger.Fatal("我是错误")

	// //等价于print();panic();
	// loger.Panic("我是错误")

	// //log的导出函数
	// //导出函数基于std,std是标准错误输出
	// //var std = New(os.Stderr, "", LstdFlags)

	// //获取输出项
	// fmt.Println(utillog.Instance().Flags())
	// //获取前缀
	// fmt.Printf(utillog.Instance().Prefix())
	// //输出内容
	// utillog.Instance().Output(2, "输出内容")
	// //格式化输出
	// utillog.Instance().Printf("第%d行 内容:%s", 22, "我是错误")
	// utillog.Instance().Fatal("我是错误")
	// utillog.Instance().Panic("我是错误")

	// // NOTE: This next line is key you have to call flag.Parse() for the command line
	// // options or "flags" that are defined in the glog module to be picked up.
	// glog.Info("hi_a")
	// flag.Lookup("logtostderr").Value.Set("true")
	// flag.Lookup("log_dir").Value.Set("/loglog")
	// flag.Lookup("v").Value.Set("10")
	// flag.Parse()
	// glog.Info("hi_b")

	// fmt.Println(flag.Lookup("log_dir").Value)

	// fmt.Println(flag.Lookup("log_dir").Value)
	// glog.V(4).Info("v4a")

	// glog.V(4).Info("v4b")

	// glog.Info("updateserver start！")

	// glog.Flush()

	// var logrusLog = logrus.New()
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	// if err == nil {
	// 	logrusLog.Out = file
	// } else {
	// 	logrusLog.Info("Failed to log to file, using default stderr")
	// }
	// logrusLog.WithFields(logrus.Fields{
	// 	"filename": "123.txt",
	// }).Info("打开文件失败")

	defer func() {

		if r := recover(); r != nil {
			utillog.Instance().Fatal(r)
			// var err error
			// //check exactly what the panic was and create error.
			// switch x := r.(type) {
			// case string:
			// 	err = errors.New(x)
			// case error:
			// 	err = x
			// default:
			// 	err = errors.New("Unknow panic")
			// }
		}
	}()

	// panic(55)

	c, err := redis.Dial("tcp", "66.112.212.102:6379")
	if err != nil {
		fmt.Println(err)
		return
	}
	//密码授权
	c.Do("AUTH", "123456")
	c.Do("SET", "a", 134)
	a, err := redis.Int(c.Do("GET", "a"))

	fmt.Println(a)

	defer c.Close()

	fmt.Println("启动服务器...")

	// var cmd string

	// for {
	// 	fmt.Println("请输入命令:")
	// 	fmt.Scan(&cmd)
	// 	//fmt.Println("你输入的是:",cmd)

	// 	switch cmd {
	// 	case "getall":
	// 		GetAll()
	// 	default:
	// 		fmt.Println("不能识别的命令")
	// 	}

	// 	fmt.Println()
	// }

	utillog.Instance().Init()
	utillog.Instance().Info("启动服务器")

	//last version

	// router := NewRouter()

	// utillog.Instance().Fatal(http.ListenAndServe(":8080", router))

	//new version
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var (
		echoEndpoint = flag.String("echo_endpoint", "localhost:50051", "endpoint of YourService")
	)
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err1 := gw.RegisterCustomerHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err1 != nil {

	}

	log.Print("Greeter gRPC Server gateway start at port 8080...")
	http.ListenAndServe(":8080", mux)
}

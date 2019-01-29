# GoLearning
UpdateServer

RN_LearnProj   https://github.com/DaiYz/react-native-mobx-init 在这位前辈的基础上修改了一个小例子准备下个版本使用

## Example List
- [updateserver](/UpdateServer)
- [grpc](/grpc/src)
## 操作 grpc
grpc-gateway  
编译 proto的时候 需要 import "google/api/annotations.proto"; 这时候有一个坑 当时我的protoc是下载二进制然后偷懒直接copy  二进制文件到go的目录下 没有配置环境变量
编译annotations的时候就会出问题 后来还是规规矩矩配置环境变量 ok protoc -I resources/customer -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis resources/customer/customer.proto --go_out=plugins=grpc:resources/customer/
编译protoc文件成功

网上的编译太费劲了  直接用下面的这个  切记 -I. 加入当前目录？ 忘记plugins=grpc 会有问题 不会生成client
protoc -I greeter_proto -I%GOPATH%\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis greeter_proto/greeter_proto.proto --go_out=plugins=grpc:greeter_proto


protoc -I. -IC:\Users\Administrator\go\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis greeter_proto/greeter_proto.proto --grpc-gateway_out=logtostderr=true:.

//解释文件？
protoc -I. -IC:\Users\Administrator\go\src\github.com\grpc-ecosystem\grpc-gateway\third_party\googleapis greeter_proto/greeter_proto.proto --swagger_out=logtostderr=true:.

编译完成之后  首先执行greeter_server.exe 然后执行 greeter_gateway.exe 就可以执行测试了
至于测试使用postman body不分选中raw 然后选中json格式传送 
{
    "name":"nidaye"
}
output
{
  "message": "Hello nidaye"
}
done！

package module_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

//因为 go 编译器默认会在当前目录中查找src目录下package。 如果没有src目录，编译器就找不到相应的package
//所以关于包的章节都放在src目录下面了

/*依赖管理
Go未解决的依赖问题
1.同一环境下，不同项目使用同一包的不同版本
2.无法管理对包的特定版本的依赖

vendor路径
随着Go 1.5release版本的发布，vendor目录被添加到除GOPAHR和GOROOT之外的依赖目录查找的解决方案
在Go 1.6之前，你需要手动的设置环境变量

查找依赖包路径的解决方案如下：
1.当前包的vendor目录
2.向上目录查找，直到找到src下的vendor目录
3.GOPATH下面查找依赖包
4.GOROOT下面查找依赖包

常用的依赖管理工具
dodep：https://github.com/tools/godep
glide：https://github.com/Masterminds/glide
dep：https://github.com/golang/dep

window安装glide教程： https://www.cnblogs.com/nickchou/p/8955180.html
安装成功后，我们把src/github.com/easierway包的源文件删除，
在go-19/module_package目录下执行
glide init
glide install
会重新下载包源文件到项目go-19/vendor目录下

遇到的问题：解决go build不去vendor下查找包的问题： https://www.jb51.net/article/202458.htm
问题总结：go 编译器默认会在当前目录中查找src目录下package。如果没有src目录，编译器就找不到相应的package


*/
func TestConcurrentMap(t *testing.T) {

	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

//运行结果
//=== RUN   TestConcurrentMap
//--- PASS: TestConcurrentMap (0.00s)
//get_remote_pack_test.go:48: 10 true
//PASS

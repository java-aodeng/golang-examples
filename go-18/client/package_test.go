package client

import (
	cm "github.com/easierway/concurrent_map" //还可以导入第三方的包，这是导入一个开源项目的包，会自动git clone到本地
	"series"                                 //这是导入自定义的包 路径，golang-examples\src\series
	"testing"
)

//package（包)
//1.基本复用模块单元
//2以首字母大写来表明可被包外代码访问
//3.代码的package可以和所在目录不一致
//4.同一目录里的Go代码package要保持一致
//5.创建包前需要配置项目GOPATH路径 配置教程 https://www.freesion.com/article/4765574656/
// 注： 配置路径后，我发现go默认是src包下面的，所以项目的包创建在了 golang-examples\src\ 路径下

func TestPackage(t *testing.T) {
	//调用golang-examples\src\series包里面的的Square函数
	t.Log(series.Square(1))

	//t.Log(series.square(1)) 运行结果：undefined: series.square 小写的函数外部不能访问

	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}

//运行结果
//=== RUN   TestPackage
//--- PASS: TestPackage (0.00s)
//package_test.go:9: 1
//PASS

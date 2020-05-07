// goscript
// 扩展脚本支持,分别运行 Lua脚本 、 JavaScript脚本 、 Tengo脚本
// 比较运算耗时
package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/d5/tengo/v2"       // Tengo native Go-VM Script
	"github.com/robertkrimen/otto" // JavaScript AST抽象语法树
	"github.com/yuin/gluamapper"   // Lua map
	"github.com/yuin/gopher-lua"   // Lua5.1 VM
)

//function add(a, b)
//return a+b
//end
var luaCode = `
function testFun(tab)
    result = {}
    result["key"] = "lua"
    --result["id"]=tab["idx"]
    result["key1"] = tonumber(tab["idx"])+1

    if(tab["user"]=="test")then
        result["title"]="good Lua"
    end
    if(tab["os"]=="ios")then
        result["url"]="http://www.google.com"
    else
        result["url"]="http://www.baidu.com"
    end
    
    return result
end
`

// Tengo script code
var tengoCode = `
//mymap := func(tab) {
	//tab := {"user":"test","os":"windows"}
	tab := {"idx":i,"user":a,"os":b }
    result := {key:"",key1:"",title:"",url:""}
	result["key"] = "Teng go"
	result["key1"] = string(int(tab["idx"])+1)
	if tab["user"] == "Fast Map" {
		result["title"] = "Tengo"
	}
	if tab["os"] != "ios" {
		result["url"] = "http://github.com/d5/tengo"
	} else {
		result["url"] = "http://www.csdn.net"
	}
	//return result
//}
//mymap(tab)
`
var tengoSumCode = `
each := func(seq, fn) {
    for x in seq { fn(x) }
}

sum := 0
mul := 1
each([a, b, c, d], func(x) {
	sum += x
	mul *= x
})
`
var showprint bool = true
var count int = 50000 //10000 //循环次数
func main() {
	dic := make(map[string]string)
	dic["user"] = "Fast Map" //"test"
	dic["os"] = "Windows"    //"ios"
	dic["version"] = "1.0"
	dic["idx"] = "1"
	//showprint = false //加速,不显示cmd信息
	//showprint = true
	start0 := time.Now()

	for i := 0; i < count; i++ {
		dic["idx"] = strconv.Itoa(i)
		if (i+1)%500 == 0 {
			showprint = true
		} else {
			showprint = false
		}
		LuaTest(dic)
	}
	tmp1 := time.Since(start0).Milliseconds() //Nanoseconds() / 1000 / 1000

	start1 := time.Now()
	for i := 0; i < count; i++ {
		dic["idx"] = strconv.Itoa(i)
		if (i+1)%500 == 0 {
			showprint = true
		} else {
			showprint = false
		}
		JsTest(dic)
	}
	tmp2 := time.Since(start1).Milliseconds() //Nanoseconds() / 1000 / 1000

	start2 := time.Now()
	for i := 0; i < count; i++ {
		dic["idx"] = strconv.Itoa(i)
		if (i+1)%500 == 0 {
			showprint = true
		} else {
			showprint = false
		}
		GoTest(dic)
	}
	tmp3 := time.Since(start2).Milliseconds() //Nanoseconds() / 1000 / 1000

	//tengoTest()
	start3 := time.Now()
	for i := 0; i < count; i++ {
		dic["idx"] = strconv.Itoa(i)
		if (i+1)%500 == 0 {
			showprint = true
		} else {
			showprint = false
		}
		tengoTest(dic)
	}
	tmp4 := time.Since(start3).Milliseconds()

	fmt.Printf("Time consuming to create  %d Maps : \n\tLua Test :\t %d ms;\n\tJS Test :\t %d ms;\n", count, tmp1, tmp2)
	fmt.Printf("\tGo Test :\t %d ms;\n", tmp3)
	fmt.Printf("\tTenGo Test :\t %d ms;\n", tmp4)
	tengoSum()
}

func LuaTest(dic map[string]string) {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(luaCode); err != nil {
		panic(err)
	}
	table := L.NewTable()
	for k, v := range dic {
		L.SetTable(table, lua.LString(k), lua.LString(v))
	}

	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("testFun"),
		NRet:    1,
		Protect: true,
	}, table); err != nil {
		panic(err)
	}
	ret := L.Get(-1) // returned value
	L.Pop(1)         // remove received value
	obj := gluamapper.ToGoValue(ret, gluamapper.Option{NameFunc: printTest})
	if !showprint {
		return
	}
	fmt.Println("Lua: ", obj)
}

func printTest(s string) string {
	return s
}

func JsTest(dic map[string]string) {
	vm := otto.New()
	v, err := vm.Run(`
function testJSFun(tab) {
    result = {}
    result["key"] = "js"
    //result["id"]=tab["idx"]
    result["key1"] = Number(tab["idx"]) + 1
    if(tab["user"]=="test"){
       result["title"]="good js"
    }
    if(tab["os"]!="ios"){
        result["url"]="http://www.google.com"
    }else{
        result["url"]="http://www.baidu.com"
    }
    
    return result
}
`)
	if err != nil {
		fmt.Println(err.Error(), v)
	}
	jsa, err := vm.ToValue(dic)
	if err != nil {
		panic(err)
	}
	result, err := vm.Call("testJSFun", nil, jsa)

	tmpR, err := result.Export()
	if !showprint {
		return
	}
	fmt.Println("js object: ", tmpR)

}

//go map
func GoTest(tab map[string]string) {
	result := make(map[string]string)

	result["key"] = "go run"
	//result["id"] = tab["idx"]
	vid, _ := strconv.Atoi(tab["idx"])
	result["key1"] = strconv.Itoa(vid + 1)
	if tab["user"] == "test" {
		result["title"] = "good Go"
	}
	if tab["os"] != "ios" {
		result["url"] = "http://www.github.com"
	} else {
		result["url"] = "http://www.jianshu.com"
	}
	if !showprint {
		return
	}

	fmt.Println("go map: ", result)
}

func tengoTest(tab map[string]string) {
	script := tengo.NewScript([]byte(tengoCode))

	// set values
	//_ = script.Add("tab", tab)
	_ = script.Add("i", tab["idx"])
	_ = script.Add("a", tab["user"])
	_ = script.Add("b", tab["os"])

	// run the script
	compiled, err := script.RunContext(context.Background())
	if err != nil {
		panic(err)
	}

	// retrieve values
	tmap := compiled.Get("result") //"mymap")
	if !showprint {
		return
	}
	fmt.Println("Tengo-map: ", tmap)
}

func tengoSum() {
	// create a new Script instance
	script := tengo.NewScript([]byte(tengoSumCode))

	// set values
	_ = script.Add("a", 1)
	_ = script.Add("b", 9)
	_ = script.Add("c", 8)
	_ = script.Add("d", 4)

	// run the script
	compiled, err := script.RunContext(context.Background())
	if err != nil {
		panic(err)
	}

	// retrieve values
	sum := compiled.Get("sum")
	mul := compiled.Get("mul")
	fmt.Println("Tengo-SUM() =", sum, mul) // "22 288"
}

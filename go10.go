/**
casbin基础模型
perm元模型


[request_definition]
r = sub, obj, act // 请求三个参数 型参

[policy_definition]
p = sub, obj, act  p策略  型参

[policy_effect]
e = some(where (p.eft == allow)) // e表明匹配终止结果之后，能否通过这个表达式得到布尔值

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act  // m是入参  r和p得对应关系 能否在p里找一条到得到一个e

PERM元模型

Polic 策略 p={sbu,obj,act,eft}
			subject(sub访问实体,谁访问的它)
			object(obj访问的资源，我要去请求什么东西)
			action(act访问方法，请求方式get..)
		三者定义完会返回一个策略结果 eft
			eft(策略结果 一般为空 默认指定allow还可以定义为deny只有两种结果)
	  策略一般存储到数据库因为会有很多
      [policy_definition]
      p=sub,obj.act-----------


Matcher  r请求 p策略
满足这个表达式条件这一条 eft会被返回回来结果进入->e影响模型规定好的返回true 或者false

角色域
role_definition

g = _,_表示以角色位基础   用户是哪个角色
g = _,_,_表示以域位基础（多商户模式） 这个用户是哪个角色，属于哪个商户域

数据库存储policy


最后一步对policy 进行增删改查

自定义比较函数


 */
package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	 gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"honnef.co/go/tools/arg"
	"strings"
)

func main()  {
	//	e, err := casbin.NewEnforcer("./model.conf","./policy.csv")

	// 适配器
	a, _ := gormadapter.NewAdapter("mysql", "root:123456@tcp(127.0.0.1:3306)/gw",true) // Your driver and data source.
	e, _ := casbin.NewEnforcer("./model.conf", a)

	sub := "alice" // 想要访问资源的用户
	obj := "data1" // 将要被访问的资源
	act := "read" // 用户对资源的操作
	added ,err:= e.AddGroupingPolicy("alice", "data2_admin") // 添加角色


	//added,err := e.AddPolicy("alice", "data2", "read")
	//fmt.Println(err)
	//fmt.Println(added)
	//filteredPolicy := e.GetFilteredPolicy(0, "data2")
	//updated, err := e.UpdatePolicy([]string{"alice", "data1", "read"}, []string{"alice", "data3", "read"})
	fmt.Println(err)
	fmt.Println(added)
	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// 处理err
		fmt.Printf("%s",err)
	}

	if ok == true {
		// 允许alice读取data1
		fmt.Println("通过")
	} else {
		// 拒绝请求，抛出异常
		fmt.Println("未通过")
	}

	// 您可以使用BatchEnforce()来批量执行一些请求
	// 这个方法返回布尔切片，此切片的索引对应于二维数组的行索引。
	// 例如results[0] 是{"alice", "data1", "read"}的结果
	//results, err := e.BatchEnforce([[] []interface{}{"alice", "data1", "read"}, {"bob", datata2", "write"}, {"jack", "data3", "read"}})
}

func testk() bool{
	return false
}
func test(key string, key2 string) bool  {
	name1 := args[0].(string)
	namde2 :=agrg[1].(string)
	return  (bool)(KeyMatch((name1,namde2)).nil
}
func KeyMatch(key1 string, key2 string) bool {
	//i := strings.Index(key2, "*")
	//if i == -1 {
	//	return key1 == key2
	//}
	//
	//if len(key1) > i {
	//	return key1[:i] == key2[:i]
	//}
	//return key1 == key2[:i]
	return key1 == key2
}
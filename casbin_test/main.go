package main

import (
    "log"
    "fmt"
    "github.com/casbin/casbin/v2"
    xormadapter "github.com/casbin/xorm-adapter/v2"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

	// 使用MySQL数据库初始化一个Xorm适配器
	a, err := xormadapter.NewAdapter("mysql", "viedo_server:123456@tcp(127.0.0.1:3306)/casbin_test", true)
	if err != nil {
	    log.Fatalf("error: adapter: %s", err)
	}

	e, err := casbin.NewEnforcer("model.conf", a)
	if err != nil {
	    log.Fatalf("error: enforcer: %s", err)
	}

	//removed, _ := e.RemovePolicy("alice", "data2")
	//fmt.Println(removed);

	/*
	rules := [][] string {
	    []string {"data2_admin", "data1", "read"},
	    []string {"data2_admin", "data2", "read"},
	}
	areRulesAdded, _ := e.AddPolicies(rules)
	fmt.Println(areRulesAdded);
	*/
	
	sub := "alice" // 想要访问资源的用户。
	obj := "data1" // 将被访问的资源。
	act := "read" // 用户对资源执行的操作。

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
	    log.Fatalf("error: enforcer: %s", err)
	}

	if ok == true {
	    // 允许alice读取data1
	    fmt.Println("ok");
	} else {
	    // 拒绝请求，抛出异常
	    fmt.Println("forbiden");
	}


}

package main

import (
	"fmt"
	"memCache/cache"
	"unsafe"
)

func main() {
	/*
		//cache := cache.NewMemCache()
		cache := cache_server.NewMemCache()
		cache.SetMaxMemory("300GB")

		cache.Set("int", 1, time.Second)
		cache.Set("bool", false, time.Second)
		cache.Set("data", map[string]interface{}{"a": 1}, time.Second)

		cache.Set("int", 1)
		cache.Set("bool", false)
		cache.Set("data", map[string]interface{}{"a": 1})

		cache.Get("int")
		cache.Del("int")
		cache.Flush()
		fmt.Println(cache.Keys())

		cache.SetMaxMemory("300GB")

		cache.Set("int", 1, time.Second)
		cache.Set("bool", false, time.Second)
		cache.Set("data", map[string]interface{}{"a": 1}, time.Second)

		cache.Set("int", 1)
		cache.Set("bool", false)
		cache.Set("data", map[string]interface{}{"a": 1})

		fmt.Println(cache.Get("int"))
		fmt.Println(cache.Get("bool"))
		fmt.Println(cache.Keys())
	*/
	cache.GetValSize(1)
	cache.GetValSize(false)
	cache.GetValSize("张三，李四，王五，赵六")
	cache.GetValSize(map[string]string{
		"name": "张三0，李四0，王五0，赵六0",
		"addr": "张三，李四，王五，赵六 张三，李四，王五，赵六",
	})

	fmt.Println(unsafe.Sizeof(1))
	fmt.Println(unsafe.Sizeof(false))
	fmt.Println(unsafe.Sizeof("张三，李四，王五，赵六"))
	fmt.Println(unsafe.Sizeof(map[string]string{
		"name": "张三0，李四0，王五0，赵六0",
		"addr": "张三，李四，王五，赵六 张三，李四，王五，赵六",
	}))

}

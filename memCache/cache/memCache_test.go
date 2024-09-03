package cache

import (
	"testing"
	"time"
)

func TestCacheOP(t *testing.T) {
	testDate := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"abc", 123, time.Second * 10},
		{"defg", false, time.Second * 11},
		{"abcdff", true, time.Second * 12},
		{"abcsafvvs", map[string]interface{}{"a": 1, "b": false}, time.Second * 13},
		{"abcsss", "dafghdtjn", time.Second * 14},
		{"abcdddd", "这是一个字符串", time.Second * 15},
	}
	c := NewMemCache()
	c.SetMaxMemory("10MB")
	for _, item := range testDate {
		c.Set(item.key, item.val, item.expire)
		val, ok := c.Get(item.key)
		if !ok {
			t.Error("缓存取值失败")
		}
		if item.key != "abcsafvvs" && val != item.val {
			t.Error("缓存取值数据与预期不一致")
		}
		_, ok1 := val.(map[string]interface{})
		if item.key == "abcsafvvs" && !ok1 {
			t.Error("缓存取值数据与预期不一致")
		}
	}
	if int64(len(testDate)) != c.Keys() {
		t.Error("缓存数量不一致")
	}
	c.Del(testDate[0].key)
	c.Del(testDate[1].key)
	if int64(len(testDate)) != c.Keys()+2 {
		t.Error("缓存数量不一致")
	}

	time.Sleep(time.Second * 16)
	if c.Keys() != 0 {
		t.Error("过期缓存清空失败")
	}
}

package limiter

import (
	"sync"
	"time"
)

type Limiter struct {
	tb *TokenBuket
}

type TokenBuket struct {
	// 锁
	mu sync.Mutex
	// 桶大小
	size int
	// 当前 token 数
	count int
	// 填充速率，即每隔多久补充一个token
	rateLimit time.Duration
	// 最后成功请求时间
	lastRequestTime time.Time
}

func (tb *TokenBuket) fillToken() {
	tb.count += tb.getFillTokenCount()
}

func (tb *TokenBuket) getFillTokenCount() int {
	if tb.count >= tb.size {
		return 0
	}
	if !tb.lastRequestTime.IsZero() {
		duration := time.Now().Sub(tb.lastRequestTime)
		count := int(duration / tb.rateLimit)
		if tb.size >= tb.count {
			return count
		} else {
			return tb.size - tb.count
		}
	}
	return 0
}

func (tb *TokenBuket) allow() bool {
	// 填充
	tb.fillToken()
	if tb.count > 0 {
		tb.count--
		tb.lastRequestTime = time.Now()
		return true
	}
	return false
}

// 初始化限流器
func NewLimiter(r time.Duration, size int) *Limiter {
	return &Limiter{
		tb: &TokenBuket{
			rateLimit: r,
			size:      size,
			count:     size,
		},
	}
}

func (l *Limiter) Allow() bool {
	l.tb.mu.Lock()
	defer l.tb.mu.Unlock()
	// 计算补充token的数
	// 当前token是否满足本次消耗
	return l.tb.allow()
}

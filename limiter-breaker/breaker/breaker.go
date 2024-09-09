package breaker

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	STATE_CLOASE = iota
	STATE_OPEN
	STATE_HALF_OPEN
)

type Breaker struct {
	mu                sync.Mutex
	state             int
	failureThreshold  int
	successThreshold  int
	halfMaxRequest    int
	halfCycleReqCount int
	timeout           time.Duration
	failureCount      int
	successCount      int
	cycleStartTime    time.Time
}

// 新建断路器
func NewBreaker(failureThreshold, successThreshold, halfMaxRequest int, timeout time.Duration) *Breaker {
	return &Breaker{
		state:            STATE_CLOASE,
		failureThreshold: failureThreshold,
		successThreshold: successThreshold,
		halfMaxRequest:   halfMaxRequest,
		timeout:          timeout,
	}
}

func (b *Breaker) Exec(f func() error) error {
	b.before()
	fmt.Printf("%+v\n", b)
	if b.state == STATE_OPEN {
		return errors.New("断路器处于打开状态无法访问服务")
	}

	if b.state == STATE_CLOASE {
		err := f()
		b.after(err)
		return err
	}

	if b.state == STATE_HALF_OPEN {
		if b.halfCycleReqCount < b.halfMaxRequest {
			err := f()
			b.after(err)
			return err
		} else {
			return errors.New("断路器处于半开闭状态，单位时间请求次数过多，请稍后再试")
		}
	}
	return nil
}

func (b *Breaker) before() {
	b.mu.Lock()
	defer b.mu.Unlock()
	switch b.state {
	case STATE_OPEN:
		// 打开状态
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			b.state = STATE_HALF_OPEN
			b.reset()
			return
		}
	case STATE_HALF_OPEN:
		// 半开半闭状态
		if b.successCount >= b.successThreshold {
			b.state = STATE_OPEN
			b.reset()
			return
		}
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			b.cycleStartTime = time.Now()
			b.halfCycleReqCount = 0
			return
		}
	case STATE_CLOASE:
		// 闭合状态(关闭状态)
		if b.cycleStartTime.Add(b.timeout).Before(time.Now()) {
			b.reset()
		}
	}
}

func (b *Breaker) after(err error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if err == nil {
		b.onSuccess()
	} else {
		b.onFailure()
	}
}

func (b *Breaker) onSuccess() {
	b.failureCount = 0
	if b.state == STATE_HALF_OPEN {
		b.successCount++
		b.halfCycleReqCount++
		if b.successCount >= b.successThreshold {
			b.state = STATE_CLOASE
			b.reset()
		}
	}
}

func (b *Breaker) onFailure() {
	b.successCount = 0
	b.failureCount++
	if b.state == STATE_HALF_OPEN || (b.state == STATE_CLOASE && b.failureCount >= b.failureThreshold) {
		b.state = STATE_OPEN
		b.reset()
		return
	}
}

func (b *Breaker) reset() {
	b.successCount = 0
	b.failureCount = 0
	b.halfCycleReqCount = 0
	b.cycleStartTime = time.Now()
}

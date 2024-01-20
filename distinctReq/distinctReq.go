package distinctReq

import "time"

// 限流
// 算法1  令牌桶算法
// 维护一个令牌桶，一恒定的速率向桶中添加令牌，
// 每个请求到达消耗一个令牌。如果桶中的令牌不够，则请求会被限制。
// 算法2   漏桶算法
// 维护一个漏桶，以恒定的速率漏水，
// 当每个请求到达，检查漏桶是否还有容量来接受请求，如果没有足够容量，则请求会被限制

// 使用场景
// 1、保护服务    请求过多的影响，防止系统过载，资源耗尽，服务奔溃
// 2、防止雪崩效应  一次性大量请求导致，保证服务稳定性
// 3、资源管理     防止请求展用过多的资源，导致其他请求访问不到
// 4、保护数据库   防止过多的查询，以致数据库性能下降
 type TokenBucket struct {
	capacity int64								// 桶的容量大小
	fillInterval time.Duration		// 填充频率
	availableTokens int64					// 可用token数量
	lastTime time.Time						// 上次填充时间
 }

 func NewTokenBucket (capacity int64, fillInterval time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity: capacity,
		fillInterval: fillInterval,
		availableTokens: capacity,
		lastTime: time.Now(),
	}
 }


 // 获取一个令牌， 桶中有可用令牌则返回true， 否则返回false
 func (b *TokenBucket) Take() bool {

	now := time.Now()
	b.availableTokens += b.toTokeAdd(now)

	if b.availableTokens > b.capacity {
		b.availableTokens = b.capacity
	}
	if b.availableTokens > 0 {
		b.availableTokens--
		return true
	}
	return false
 }

 func (b *TokenBucket) toTokeAdd(now time.Time) int64 {
		d := now.Sub(b.lastTime)
		if d < 0 {
			return 0 
		}

		fillNum := int64(d / b.fillInterval)

		b.lastTime = b.lastTime.Add(time.Duration(fillNum) * b.fillInterval)
		return fillNum
 }



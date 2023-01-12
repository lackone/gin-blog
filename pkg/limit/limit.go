package limit

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimitIface interface {
	//获取对应限流器的键值对名称
	Key(c *gin.Context) string
	//获取令牌桶
	GetBucket(key string) (*ratelimit.Bucket, bool)
	//添加多个令牌桶
	AddBuckets(rules ...LimitBucketRule) LimitIface
}

type LimitBucketRule struct {
	Key          string        //自定义键值对名称
	FillInterval time.Duration //间隔多久时间放N个令牌
	Capacity     int64         //令牌桶的容量
	Quantum      int64         //每次到达间隔时间后所放的具体令牌数量
}

type Limit struct {
	limitBuckets map[string]*ratelimit.Bucket
}

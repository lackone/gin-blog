package limit

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimit struct {
	*Limit
}

func NewMethodLimit() LimitIface {
	return &MethodLimit{
		Limit: &Limit{
			limitBuckets: make(map[string]*ratelimit.Bucket),
		},
	}
}

func (m *MethodLimit) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (m *MethodLimit) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := m.limitBuckets[key]
	return bucket, ok
}

func (m *MethodLimit) AddBuckets(rules ...LimitBucketRule) LimitIface {
	for _, rule := range rules {
		if _, ok := m.limitBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
			m.limitBuckets[rule.Key] = bucket
		}
	}
	return m
}

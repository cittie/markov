package markov

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus_Rand(t *testing.T) {
	s := &Status{ratios: []int64{1, 2, 3, 4}}

	t.Run("Ratios", func(t *testing.T) {
		counts := make(map[int]int)

		for i := 0; i < 10000; i++ {
			counts[s.Rand()] += 1
		}

		// 误差小于100
		for k, v := range counts {
			r := (k+1)*1000 - v
			assert.True(t, r < 100 || r > -100)
		}
	})
}

func TestChain_Next(t *testing.T) {
	chain := &Chain{
		curStatus: 0,
		status: []IStatus{
			&Status{ratios: []int64{0, 4, 4, 9}},
			&Status{ratios: []int64{1, 0, 4, 9}},
			&Status{ratios: []int64{1, 4, 0, 9}},
			&Status{ratios: []int64{1, 4, 4, 0}},
		},
	}

	// 上面设置了所有条件矩阵都不能跳回自身
	for i := 0; i < 1000; i++ {
		assert.True(t, chain.CurStatusIdx() != chain.Next())
	}
}

func BenchmarkChain_Next(b *testing.B) {
	// 用16*16的矩阵进行性能测试
	r := make([]int64, 16)
	for i := int64(0); i < 16; i++ {
		r[i] = i
	}

	status := make([]IStatus, 16)
	for i := 0; i < 16; i++ {
		status[i] = &Status{ratios: r}
	}

	chain := &Chain{
		curStatus: 0,
		status:    status,
	}

	for i := 0; i < b.N; i++ {
		chain.Next()
	}
}

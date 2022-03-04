package markov

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatus_Rand(t *testing.T) {
	s := &Status{ratios: []int64{1, 2, 3, 4}}

	t.Run("Ratios", func(t *testing.T) {
		counts := make(map[int]int, 4)

		for i := 0; i < 10000; i++ {
			counts[s.Rand()] += 1
		}

		// 误差小于100
		for k, v := range counts {
			r := (k+1)*1000 - v
			assert.True(t, r < 100 || r > -100)
		}

		t.Logf("counts: %v", counts)
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

func TestChainCreate(t *testing.T) {
	ch := NewChain()
	assert.Nil(t, ch.AddStatus(NewStatus("line1", []int64{1}), []int64{}))
	assert.Nil(t, ch.AddStatus(NewStatus("line2", []int64{1, 1}), []int64{2}))
	assert.Nil(t, ch.AddStatus(NewStatus("line3", []int64{0, 3, 3}), []int64{4, 5}))
	assert.Nil(t, ch.AddStatus(NewStatus("line4", []int64{4, 3, 2, 1}), []int64{9, 6, 3}))

	t.Logf("chain status: %v", ch.ListStatus())

	assert.Nil(t, ch.SetCurStatusIdx(0))

	for i := 0; i < 20; i++ {
		n := ch.Next()
		fmt.Println(n)
	}

	fmt.Println(ch.CurStatusIdx())
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

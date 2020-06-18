package markov

import (
	"math/rand"
)

/*
	概率用整数表示
*/

type Chain struct {
	curStatus int
	status    []IStatus
}

type Status struct {
	ratios []int64
}

// Chain

// NewChain
func NewChain() *Chain {
	return &Chain{
		curStatus: -1,
		status:    make([]IStatus, 0, 16),
	}
}

func (mc *Chain) ListStatus() []IStatus {
	return mc.status
}

func (mc *Chain) GetStatus(idx int) (IStatus, error) {
	if idx < 0 || idx >= len(mc.status) {
		return nil, ErrIdxOutOfRange
	}
	return mc.status[idx], nil
}

func (mc *Chain) AddStatus(status IStatus) error {
	if status.Len() != len(mc.status)+1 {
		return ErrInvalidStatus
	}

	mc.status = append(mc.status, status)

	for _, s := range mc.status {
		s.Increase()
	}

	return nil
}

func (mc *Chain) UpdateStatus(idx int, status IStatus) error {
	if idx < 0 || idx >= len(mc.status) {
		return ErrIdxOutOfRange
	}

	if status.Len() != len(mc.status) {
		return ErrInvalidStatus
	}

	mc.status[idx] = status

	return nil
}

func (mc *Chain) DelStatus(idx int) {
	if idx < 0 || idx >= len(mc.status) {
		return
	}

	// 删除整行
	if mc.curStatus == idx {
		mc.curStatus = -1
	}

	mc.status = append(mc.status[:idx], mc.status[idx+1:]...)

	// 删除每行对应的部分
	for i := 0; i < len(mc.status); i++ {
		mc.status[i].Del(idx)
	}
}

func (mc *Chain) CurStatusIdx() int {
	return mc.curStatus
}

func (mc *Chain) SetCurStatusIdx(idx int) error {
	if idx < 0 || idx >= len(mc.status) {
		return ErrIdxOutOfRange
	}

	mc.curStatus = idx

	return nil
}

func (mc *Chain) Next() int {
	if mc.curStatus == -1 {
		return -1
	}

	mc.curStatus = mc.status[mc.curStatus].Rand()
	return mc.curStatus
}

// Status
func (s *Status) Len() int {
	return len(s.ratios)
}

func (s *Status) Del(idx int) {
	if idx < 0 || idx >= len(s.ratios) {
		return
	}

	s.ratios = append(s.ratios[:idx], s.ratios[idx+1:]...)
}

func (s *Status) ShowStatus() string {
	return ""
}

// Rand
// 按权重随机出当前分组数量
func (s *Status) Rand() int {
	if len(s.ratios) == 0 {
		return -1
	}

	var total int64
	for i := 0; i < len(s.ratios); i++ {
		total += s.ratios[i]
	}

	n := rand.Int63n(total)

	for i := 0; i < len(s.ratios); i++ {
		if n < s.ratios[i] {
			return i
		}
		n -= s.ratios[i]
	}

	return len(s.ratios)
}

func (s *Status) Increase() {
	s.ratios = append(s.ratios, 0)
}

// SetRatio 直接赋值
func (s *Status) SetRatio(ratios []int64) {
	s.ratios = ratios
}

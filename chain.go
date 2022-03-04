package markov

import (
	"fmt"
)

/*
	不支持并发
	概率用整数表示
*/

type Chain struct {
	curStatus int
	status    []IStatus
}

type Status struct {
	name   string
	ratios []int64
}

func NewChain() *Chain {
	return &Chain{
		curStatus: StatusNotAvailable,
		status:    make([]IStatus, 0, 16),
	}
}

// Chain 相关接口实现

func (mc *Chain) SetStatus(status []IStatus) error {
	for _, s := range status {
		if s.Len() != len(status) {
			return ErrInvalidStatus
		}
	}

	mc.status = status

	return nil
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

func (mc *Chain) AddStatus(status IStatus, rates []int64) error {
	if status.Len() != len(mc.status)+1 {
		return ErrInvalidStatus
	}

	if len(rates) != len(mc.status) {
		return ErrInvalidRates
	}

	for i, s := range mc.status {
		s.AddRate(rates[i])
	}

	mc.status = append(mc.status, status)

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
		mc.curStatus = StatusNotAvailable
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
	if mc.curStatus == StatusNotAvailable {
		return StatusNotAvailable
	}

	mc.curStatus = mc.status[mc.curStatus].Rand()
	return mc.curStatus
}

// Status 相关接口实现

func NewStatus(name string, ratio []int64) *Status {
	return &Status{
		name:   name,
		ratios: ratio,
	}
}

func (s Status) String() string {
	return fmt.Sprintf("[name:%s, ratios:%v]", s.name, s.ratios)
}

func (s Status) GetName() string {
	return s.name
}

func (s Status) Len() int {
	return len(s.ratios)
}

func (s *Status) Del(idx int) {
	if idx < 0 || idx >= len(s.ratios) {
		return
	}

	s.ratios = append(s.ratios[:idx], s.ratios[idx+1:]...)
}

func (s Status) ShowRation() string {
	return fmt.Sprintf("%v", s.ratios)
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

	n := Int63n(total)

	for i := 0; i < len(s.ratios); i++ {
		if n < s.ratios[i] {
			return i
		}
		n -= s.ratios[i]
	}

	return len(s.ratios)
}

func (s *Status) AddRate(rate int64) {
	s.ratios = append(s.ratios, rate)
}

// SetRatio 直接赋值
func (s *Status) SetRatio(ratios []int64) {
	s.ratios = ratios
}

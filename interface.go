package markov

/*
	马尔科夫链
*/

type IChain interface {
	// SetStatus 直接将矩阵赋值
	SetStatus(status []IStatus) error

	// ListStatus 列出当前所有状态
	ListStatus() []IStatus

	// GetStatus 获取当前状态
	// 如果当前尚未进行过随机过程或当前随机过程已被删除，返回-1
	GetStatus(idx int) (IStatus, error)

	// AddStatus 添加状态
	// status 长度需要为当前已有status长度+1
	// rates 长度需要和当前已有status长度一致
	AddStatus(status IStatus, rates []int64) error

	// UpdateStatus 替换状态
	UpdateStatus(idx int, status IStatus) error

	// DelStatus 删除状态
	// 如果删除的是当前状态，需要再指定一个开始值
	DelStatus(idx int)

	// CurStatusIdx 当前状态序号
	CurStatusIdx() int

	// SetCurStatusIdx 设置当前状态序号
	SetCurStatusIdx(idx int) error

	// Next 跳转到下个状态，返回跳转后的值
	// 开始跳转前需要指定一个值
	// 如果之前未指定开始状态，会返回-1
	Next() int

	// TODO 获取概率分布
}

type IStatus interface {
	GetName() string

	// Len 长度
	Len() int

	// AddRate 增加新概率
	AddRate(rate int64)

	// Del 删掉第n位
	Del(idx int)

	// ShowRation 显示当前概率
	ShowRation() string

	// Rand 随机生成下次状态
	// 如果当前概率列表为空，返回-1，
	Rand() int
}

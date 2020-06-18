package markov

/*
	马尔科夫链
*/

type IChain interface {
	// ListStatus 列出当前所有状态
	ListStatus() []IStatus

	// 获取当前状态
	// 如果当前尚未进行过随机过程或当前随机过程已被删除，返回-1
	GetStatus(idx int) (IStatus, error)

	// 添加状态
	AddStatus(status IStatus) error

	// 替换状态
	UpdateStatus(idx int, status IStatus) error

	// 删除状态
	DelStatus(idx int)

	// 当前状态序号
	CurStatusIdx() int

	// 设置当前状态序号
	SetCurStatusIdx(idx int) error

	// 跳转到下个状态，返回跳转后的值
	Next() int

	// TODO 获取概率分布
}

type IStatus interface {
	// 长度
	Len() int

	// 自增1位
	Increase()

	// 删掉第n位
	Del(idx int)

	// 显示当前概率
	ShowStatus() string

	// 随机生成下次状态
	// 如果当前概率列表为空，返回-1，
	Rand() int
}

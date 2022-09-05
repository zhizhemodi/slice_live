package newsudoku

type node struct {
	num       uint8
	posi      []uint8
	changable bool
}

type step struct {
	id    int
	x     int
	y     int
	value uint8
}

type TablePlace struct {
	table       [9][9]node // 九宫格桌
	stepList    []step     // 步骤记录
	Blank       int        // 剩余空格
	minPosiStep step       // 最小可能步
}

package newsudoku

import "fmt"

func (tp *TablePlace) put(k1, k2 int, value uint8) error {
	if !tp.table[k1][k2].changable {
		return fmt.Errorf("该格无法修改")
	}
	tp.table[k1][k2].num = value
	return nil
}

func (tp *TablePlace) delete(k1, k2 int) {
	tp.table[k1][k2].num = 0
}

func (tp *TablePlace) FindPosi(k1, k2 int) ([]uint8, error) {
	if tp.table[k1][k2].num != 0 {
		return nil, fmt.Errorf("该格已有数值")
	}
	sto := make(map[uint8]bool)
	for i := 0; i < 9; i++ {
		if tp.table[k1][i].num != 0 {
			sto[tp.table[k1][i].num] = true
		}
		if tp.table[i][k2].num != 0 {
			sto[tp.table[i][k2].num] = true
		}
	}
	for i := k1 / 3 * 3; i < (k1/3+1)*3; i++ {
		for j := k2 / 3 * 3; j < (k2/3+1)*3; j++ {
			if tp.table[i][j].num != 0 {
				sto[tp.table[i][j].num] = true
			}
		}
	}
	re := []uint8{}
	for i := 1; i <= 9; i++ {
		if !sto[uint8(i)] {
			re = append(re, uint8(i))
		}
	}
	return re, nil
}

// 单线程 深度优先算法的回溯部分
func (tp *TablePlace) backToFork() {
	for {
		step := tp.stepList[len(tp.stepList)-1]
		if len(tp.table[step.x][step.y].posi) > 1 {
			tp.table[step.x][step.y].posi = tp.table[step.x][step.y].posi[1:]
			tp.table[step.x][step.y].num = tp.table[step.x][step.y].posi[0]
			return
		}
		tp.table[step.x][step.y].num = 0
		tp.stepList = tp.stepList[:len(tp.stepList)-1]
		tp.Blank++
	}
}

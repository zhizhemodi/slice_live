package newsudoku

import "fmt"

func New() TablePlace {
	return TablePlace{}
}

func (tp *TablePlace) Init(data [9][9]uint8) {
	blank := 0
	for k1, v1 := range data {
		for k2, v2 := range v1 {
			tp.table[k1][k2].num = v2
			if v2 == 0 {
				tp.table[k1][k2].changable = true
				blank++
			} else {
				tp.table[k1][k2].changable = false
			}
		}
	}
	tp.Blank = blank
}

func (tp *TablePlace) PrintTable() {
	for _, v1 := range tp.table {
		fmt.Print("|")
		for _, v2 := range v1 {
			fmt.Printf(" %v", v2)
		}
		fmt.Println("|")
	}
	fmt.Println(tp.stepList)
}

func (tp *TablePlace) Next() {
	tp.minPosiStep = step{value: 9}
	for k1, v1 := range tp.table {
		for k2, v2 := range v1 {
			if v2.num == 0 {
				list, err := tp.FindPosi(k1, k2)
				if err != nil {
					return
				}
				tp.table[k1][k2].posi = list
				if len(list) == 0 {
					tp.backToFork()
					return
				}
				if len(list) == 1 {
					tp.table[k1][k2].num = list[0]
					tp.stepList = append(tp.stepList, step{
						id:    len(tp.stepList) + 1,
						x:     k1,
						y:     k2,
						value: list[0],
					})
				}
				if uint8(len(list)) < tp.minPosiStep.value {
					tp.minPosiStep = step{
						x:     k1,
						y:     k2,
						value: uint8(len(list)),
					}
				}
			}
		}
	}
	if tp.minPosiStep.value >= 2 {
		tp.table[tp.minPosiStep.x][tp.minPosiStep.y].num = tp.table[tp.minPosiStep.x][tp.minPosiStep.y].posi[0]
		tp.stepList = append(tp.stepList, step{
			id:    len(tp.stepList) + 1,
			x:     tp.minPosiStep.x,
			y:     tp.minPosiStep.y,
			value: tp.table[tp.minPosiStep.x][tp.minPosiStep.y].posi[0],
		})
		tp.Blank--
	}
}

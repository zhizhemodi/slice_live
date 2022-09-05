package newsudoku

import (
	"fmt"
	"os"
)

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

func (tp *TablePlace) PrintTable() error {
	f, err := os.Create("answer.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	for _, v1 := range tp.table {
		fmt.Fprint(f, "|")
		for _, v2 := range v1 {
			fmt.Fprintf(f, " %v", v2.num)
		}
		fmt.Fprintln(f, "|")
	}
	fmt.Println(" ")
	return nil
}

func (tp TablePlace) Next(downstream chan string, id string, posi int) {
	for {
		if posi > 0 {
			tp.table[tp.minPosiStep.x][tp.minPosiStep.y].num = tp.table[tp.minPosiStep.x][tp.minPosiStep.y].posi[posi]
			tp.stepList = append(tp.stepList, step{
				id:    len(tp.stepList) + 1,
				x:     tp.minPosiStep.x,
				y:     tp.minPosiStep.y,
				value: tp.table[tp.minPosiStep.x][tp.minPosiStep.y].posi[posi],
			})
			tp.Blank--
			posi = -1
		}
		err := tp.once(downstream, id)
		if err != nil {
			return
		}
		if tp.Blank == 0 {
			tp.PrintTable()
			downstream <- id + " success"
			break
		}
		if tp.minPosiStep.value >= 2 {
			for i := 1; i < int(tp.minPosiStep.value); i++ {
				go tp.Next(downstream, fmt.Sprintf("%v%v", id, i), i)
			}
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
}

func (tp *TablePlace) once(downstream chan string, id string) error {
	tp.minPosiStep = step{value: 9}
	for k1, v1 := range tp.table {
		for k2, v2 := range v1 {
			if v2.num == 0 {
				list, err := tp.FindPosi(k1, k2)
				if err != nil {
					return fmt.Errorf("fail")
				}
				tp.table[k1][k2].posi = list
				if len(list) == 0 {
					return fmt.Errorf("fail")
				}
				if len(list) == 1 {
					tp.table[k1][k2].num = list[0]
					tp.stepList = append(tp.stepList, step{
						id:    len(tp.stepList) + 1,
						x:     k1,
						y:     k2,
						value: list[0],
					})
					tp.Blank--
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
	return nil
}

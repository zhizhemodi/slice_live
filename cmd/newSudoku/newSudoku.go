package main

import (
	"fmt"
	newsudoku "slice_live/newSudoku"
	"strings"
	"time"
)

func main() {
	start := time.Now().Nanosecond()
	tp := newsudoku.New()
	data := [9][9]uint8{
		// 最难
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
		// 随便找的
		// {0,0,0,6,0,5,0,0,0},
		// {0,0,0,0,7,0,0,0,0},
		// {0,8,3,0,0,0,0,6,0},
		// {0,3,0,0,9,0,0,0,4},
		// {0,0,0,0,0,4,7,5,0},
		// {4,0,0,8,3,0,1,0,0},
		// {0,0,0,0,0,9,0,0,0},
		// {0,9,4,0,0,0,2,8,0},
		// {0,2,6,0,0,3,0,0,0},
	}
	tp.Init(data)
	c0 := make(chan string)
	tp.Next(c0, "1", -1)
	for {
		v := <-c0
		if strings.Contains(v, "success") {
			fmt.Printf("v: %v\n", v)
			break
		}
	}
	end := time.Now().Nanosecond()
	fmt.Println(end - start)
}

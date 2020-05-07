package main

import tool "github.com/kilingzhang/leetcode"

func mincostTickets(days []int, costs []int) int {
	year:=make([]int,days[len(days)-1]+1)
	for k,v:=range days{
		if k!=0{
			for i:=days[k-1]+1;i<v;i++{
				year[i]=year[days[k-1]]
			}
		}
		temp1:=min(costs[0]+year[max(v-1,0)],costs[1]+year[max(v-7,0)])
		year[v]=min(temp1,costs[2]+year[max(v-30,0)])
	}
	return year[days[len(days)-1]]
}
func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func max(i int, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func main() {
	tool.Dump(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}

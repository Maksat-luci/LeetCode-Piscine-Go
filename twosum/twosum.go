package main

import "fmt"


func TwoSum(nums []int, target int) []int {
	var integer []int
	for i :=0; i != len(nums); i++ {
		for j := i+1; j != len(nums);j++{
			if nums[i] + nums[j] == target {
				integer = append(integer, i,j)
			}
		} 
	}
	res := []int{}
	if integer != nil {
		if len(integer) > 2 {
			res = append(res, integer[0])
			res = append(res, integer[1])
			return res
		}
		return integer
	}
	return nil
}

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}

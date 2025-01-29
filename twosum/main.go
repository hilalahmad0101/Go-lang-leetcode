package main

import "fmt"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	nums := []int{3, 2, 4}
	target := 6
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			target2 := nums[i] + nums[j]
			if target == target2 {
				fmt.Println(target2)
			}
		}
	}
}

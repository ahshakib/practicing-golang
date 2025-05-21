package main

import "fmt"


func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var time string
		fmt.Scan(&time)
		hours := int(time[0]-'0')*10 + int(time[1]-'0')
		minutes := int(time[3]-'0')*10 + int(time[4]-'0')
		
		if hours == 0 {
			fmt.Printf("12:%02d AM\n", minutes)
		} else if hours < 12 {
			fmt.Printf("%02d:%02d AM\n", hours, minutes)
		} else {
			if hours == 12 {
				fmt.Printf("%02d:%02d PM\n", hours, minutes)
			} else {
				fmt.Printf("%02d:%02d PM\n", hours-12, minutes)
			}
		}
	}
}
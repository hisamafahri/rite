package helper

import "fmt"

func CheckErr(e error) {
	if e != nil {
		fmt.Println("ERROR DETECTED: ", e)
	}
}

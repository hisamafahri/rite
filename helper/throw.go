package helper

import "fmt"

func ThrowInfo(info string) {
	if info != "" {
		fmt.Println("INFO: ", info)
	}
}

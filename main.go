package main

import (
	"fmt"
	"ismartTest/simulation"
)

func main() {

	var result map[string]string
	simulation.Run("", "", "", &result)
	fmt.Println("任务处理完成", result)

}

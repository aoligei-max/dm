package main

import (
	"github.com/aoligei-max/dm"
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	dmsoft.SetDllPathW(dir+"\\dm.dll", 0)
	dm, err := dmsoft.New()
	if err != nil {
		panic(err)
	}
	defer dm.Release()

	fmt.Printf("Version: %s\n", dm.Ver())
}

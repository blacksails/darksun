package main

import (
	"fmt"
	"sync"

	"github.com/blacksails/darksun"
)

func runModules(modules []darksun.Module, dark bool) {
	var wg sync.WaitGroup
	for _, m := range modules {
		wg.Add(1)
		go func(module darksun.Module) {
			defer wg.Done()
			var err error
			if dark {
				err = module.Dark()
			} else {
				err = module.Sun()
			}
			if err != nil {
				fmt.Println(err)
			}
		}(m)
	}
	wg.Wait()
}

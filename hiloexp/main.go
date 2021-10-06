package main

import (
	"fmt"
	"hiloexp/hilo"
	"sort"
	"sync"
)

func main() {
	gen := hilo.NewHiLoGen(0, -10, 1)
	keys := make([]int64, 0)
	mux := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		token := gen.NewToken()
		wg.Add(1)
		go func(t *hilo.HiLoToken) {
			for j := 0; j < 5; j++ {
				key := gen.Key(token)
				mux.Lock()
				keys = append(keys, key)
				mux.Unlock()
			}
			wg.Done()
		}(token)
	}
	wg.Wait()
	sort.Slice(keys, func(x, y int) bool { return keys[x] < keys[y] })
	for _, key := range keys {
		fmt.Println(key)
	}
}

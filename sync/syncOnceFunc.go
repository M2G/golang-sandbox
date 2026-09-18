package sync

import (
	"fmt"
	"sync"
)

var setup = sync.OnceFunc(func() {
	fmt.Println("configuration chargée une seule fois")
})

func SyncOnceFunc() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			setup() // appelable partout, s'exécute une seule fois au total
		}()
	}
	wg.Wait()
}

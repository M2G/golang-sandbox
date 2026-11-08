package goroutine

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Say(phrase string) {
	for _, word := range strings.Fields(phrase) {
		fmt.Println(word)
		dur := time.Duration(100+rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

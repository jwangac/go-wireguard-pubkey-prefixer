package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func generate() (private, public string) {
	priv, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return "", ""
	}
	pub := priv.PublicKey()
	return priv.String(), pub.String()
}

func search(prefix string) (private, public string) {
	num := math.Pow(64, float64(len(prefix)))

	var wg sync.WaitGroup
	result := make(chan [2]string)
	var counter uint64

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				priv, pub := generate()
				atomic.AddUint64(&counter, 1)
				if strings.HasPrefix(pub, prefix) {
					result <- [2]string{priv, pub}
					return
				}
			}
		}()
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fmt.Printf("Keys generated: %d (Estimate: %.1f%%)\n", atomic.LoadUint64(&counter), float64(atomic.LoadUint64(&counter))/num*100)
		}
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		return res[0], res[1]
	}

	return "", ""
}

func main() {
	prefix := os.Args[1]
	fmt.Println(search(prefix))
}

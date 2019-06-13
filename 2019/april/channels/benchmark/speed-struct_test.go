package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

var s string
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	rand.Seed(time.Now().UnixNano())
	s = randStringRunes(1000000)
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func BenchmarkSpeed(b *testing.B) {
	cap := 100
	ch := make(chan string, cap)
	wg := &sync.WaitGroup{}
	go func() {
		defer wg.Done()
		for range ch {
		}
	}()
	b.SetBytes(1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ch <- s
	}

}

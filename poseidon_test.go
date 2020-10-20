package poseidon

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func BenchmarkNewLocalStorage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewLocalStorage(KB)
	}
}

func TestLocalStorage_SetItem(t *testing.T) {
	l := NewLocalStorage(MB)
	for i := 0; i < 100000; i++ {
		l.SetItem(strconv.Itoa(i), strconv.Itoa(i)+"abcdefghijklmnopqrstuvwxyz")
	}

	for i := 100000; i >= 0; i-- {
		v, bo := l.GetItem(strconv.Itoa(i))
		fmt.Println(i, v, bo)
	}
}

func TestLocalStorage_RemoveItem(t *testing.T) {
	l := NewLocalStorage(2000 * Bytes)
	l.SetItem("key", "abcdefghijklmnopqrstuvwxyz")
	l.SetItem("key1", "abcdefghijklmnopqrstuvwxyz1")
	time.Sleep(1 * time.Second)

	fmt.Println(l.GetItem("key"))
	fmt.Println(l.GetItem("key1"))
	l.RemoveItem("key")

	fmt.Println(l.GetItem("key"))
}

func BenchmarkLocalStorage_GetItem(b *testing.B) {
	l := NewLocalStorage(KB)
	for i := 0; i < 1000; i++ {
		l.SetItem(strconv.Itoa(i), strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.GetItem(strconv.Itoa(i))
	}
}

func BenchmarkLocalStorage_SetItem(b *testing.B) {
	l := NewLocalStorage(KB)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.SetItem(strconv.Itoa(i), strconv.Itoa(i))
	}
}

func BenchmarkLocalStorage_RemoveItem(b *testing.B) {
	l := NewLocalStorage(MB)
	for i := 0; i < 10000; i++ {
		l.SetItem(strconv.Itoa(i), strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.RemoveItem(strconv.Itoa(i))
	}
}

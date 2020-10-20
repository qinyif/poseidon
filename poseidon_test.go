package poseidon

import (
	"fmt"
	"strconv"
	"testing"
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
	l := NewLocalStorage(KB)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.RemoveItem(strconv.Itoa(i))
	}
}

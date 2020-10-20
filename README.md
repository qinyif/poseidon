# poseidon
Local LRU Cache. Is safe for concurrent use by multiple goroutines without additional locking or coordination.

## Installation

go get -u github.com/byteconv/poseidon

## Quick start

```
//KB MB GB
l := NewLocalStorage(poseidon.KB)

l.SetItem("key", "value")
l.GetItem("key")
l.RemoveItem("key")
```
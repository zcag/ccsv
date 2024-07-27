package util

import (
	"hash/fnv"
)

func Hash(val string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(val))
	return h.Sum32()
}

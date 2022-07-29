package avl

import (
	"math/rand"
	"os"
	"testing"
	"time"
)

const (
	// half-interval boundaries
	a = 1 << 20 // 2^20
	b = 1 << 22 // 2^22

	// constants from the theoretical estimate of the maximum height
	A = 1.44
	B = 0.328
)

func randIntSlice(n int) []int {
	data := make([]int, n)
	for i := range data {
		data[i] = rand.Int()
	}
	return data
}

func randStringSlice(n int) []string {
	data := make([]string, n)
	for i := range data {
		buf := make([]byte, 16)
		rand.Read(buf)
		data[i] = string(buf)
	}
	return data
}

var (
	dataSetsInt    [][]int
	dataSetsString [][]string
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())
	dataSetsInt = [][]int{
		randIntSlice(1 << 5),
		randIntSlice(1 << 10),
		randIntSlice(1 << 15),
	}
	dataSetsString = [][]string{
		randStringSlice(1 << 5),
		randStringSlice(1 << 10),
		randStringSlice(1 << 15),
	}
	os.Exit(m.Run())
}

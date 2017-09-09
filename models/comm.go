package models

import (
	"math/rand"
	"time"
)

func Random() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100; i++ {
		nu := r.Intn(20000)
		if nu > 10000 && nu < 65535 {
			return nu

		}

	}

	return 20000
}

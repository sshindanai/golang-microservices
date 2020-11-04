package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBestBoxNotFit(t *testing.T) {
	var availableBoxes = []Box{
		Box{ID: 1, Width: 10, Height: 10, Length: 20, MaxWeight: 50},
		Box{ID: 2, Width: 10, Height: 20, Length: 20, MaxWeight: 50},
		Box{ID: 3, Width: 15, Height: 20, Length: 25, MaxWeight: 80},
		Box{ID: 4, Width: 15, Height: 30, Length: 50, MaxWeight: 100},
		Box{ID: 5, Width: 30, Height: 30, Length: 60, MaxWeight: 120},
		Box{ID: 6, Width: 40, Height: 40, Length: 40, MaxWeight: 120},
		Box{ID: 7, Width: 50, Height: 40, Length: 45, MaxWeight: 120},
		Box{ID: 8, Width: 60, Height: 60, Length: 50, MaxWeight: 150},
	}

	product := Product{Name: "mango", Width: 1160, Height: 10, Length: 10, Weight: 50}
	output := make(chan Box)
	go getBestBox(availableBoxes, product, output)

	result := <-output
	assert.NotNil(t, result)
}

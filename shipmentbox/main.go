package main

import (
	"fmt"
	"sync"
)

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

type Product struct {
	Name                          string
	Width, Height, Length, Weight int //Millimetre
}

type Box struct {
	ID, Width, Height, Length, MaxWeight int //Millimetre
}

func main() {
	products := []Product{
		Product{Name: "PS 5", Width: 60, Height: 10, Length: 10, Weight: 50},
		Product{Name: "mango", Width: 60, Height: 10, Length: 10, Weight: 50},
		Product{Name: "mango2", Width: 24, Height: 29, Length: 10, Weight: 110},
	}

	output := make(chan Box)
	var wg sync.WaitGroup
	for _, product := range products {
		wg.Add(1)
		go getBestBox(availableBoxes, product, output, &wg)
		results := <-output

		fmt.Println("The product:", product.Name)
		fmt.Println("The box for this product: ID", results.ID)
	}
	wg.Wait()
	close(output)
}

func getBestBox(availableBoxes []Box, product Product, output chan Box, wg *sync.WaitGroup) {

	for _, box := range availableBoxes {
		if product.Weight > box.MaxWeight || product.Width > box.Width || product.Height > box.Width || product.Length > box.Length {
			continue
		} else {
			output <- box
		}
		wg.Done()
	}
	return
}

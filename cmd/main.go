package main

import (
	"fmt"

	"github.com/branotix/local/26_02_2026/pulock"
)

func main() {
	var temno *pulock.Pulock = pulock.NewPulock("Pulock Kumar", "Male", "Unmeried", "017730588885", "cryptohashlock@gmail.com", "Hindu", 23, 68)

	fmt.Println(*temno)
}

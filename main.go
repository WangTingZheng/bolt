package main

import (
	"fmt"
)

func main() {
	db := DB()
	err := db.Open("db", 0666)
	if err != nil {
		fmt.Print(err)
	}

	t, terr := db.RWTransaction()
	if terr != nil {
		fmt.Print(terr)
	}

	perr := t.Put("widgets", []byte("foo"), []byte("bar"))
	if perr != nil {
		fmt.Print(perr)
	}
	value := t.Get("widgets", []byte("foo"))

	fmt.Print(value)
}

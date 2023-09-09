package testgopkg

import (
	"fmt"
)

type Retvain struct {
	word string
}

func New(word string) *Retvain {
	return &Retvain{word: word}
}

func (r *Retvain) Say() {
	fmt.Println("Retvain is v2.0.0 say " + r.word)
}

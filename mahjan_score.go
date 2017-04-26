package main

import (
	"fmt"

	"github.com/hasokon/mahjan"
)

// MahjanScore is for score
type MahjanScore struct {
	person mahjan.Person
	tsumo  bool
	hu     uint
	han    uint
}

func (ms MahjanScore) getMahjanScore() string {
	m := mahjan.New()
	return m.Score(ms.hu, ms.han, ms.person, ms.tsumo)
}

func (ms MahjanScore) String() string {
	p := "親"
	if ms.person == mahjan.Child {
		p = "子"
	}

	t := "ロン"
	if ms.tsumo {
		t = "ツモ"
	}

	return fmt.Sprintf("%s %s %d符%d翻", p, t, ms.hu, ms.han)
}

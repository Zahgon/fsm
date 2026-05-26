//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door { _ = "STUB: not implemented"; return nil }

func (d *Door) enterState(e *fsm.Event) { _ = "STUB: not implemented"; return }

func main() {
	door := NewDoor("heaven")

	err := door.FSM.Event(context.Background(), "open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event(context.Background(), "close")
	if err != nil {
		fmt.Println(err)
	}
}

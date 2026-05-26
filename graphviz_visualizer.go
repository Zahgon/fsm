package fsm

import (
	"bytes"
)

// Visualize outputs a visualization of a FSM in Graphviz format.
func Visualize(fsm *FSM) string { _ = "STUB: not implemented"; return "" }

// we sort the key alphabetically to have a reproducible graph output

func writeHeaderLine(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func writeTransitions(buf *bytes.Buffer, sortedEKeys []eKey, transitions map[eKey]string) {
	_ = "STUB: not implemented"
	return
}

func writeStates(buf *bytes.Buffer, current string, sortedStateKeys []string) {
	_ = "STUB: not implemented"
	return
}

func writeFooter(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

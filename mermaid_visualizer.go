package fsm

import (
	"bytes"
)

const highlightingColor = "#00AA00"

// MermaidDiagramType the type of the mermaid diagram type
type MermaidDiagramType string

const (
	// FlowChart the diagram type for output in flowchart style (https://mermaid-js.github.io/mermaid/#/flowchart) (including current state)
	FlowChart MermaidDiagramType = "flowChart"
	// StateDiagram the diagram type for output in stateDiagram style (https://mermaid-js.github.io/mermaid/#/stateDiagram)
	StateDiagram MermaidDiagramType = "stateDiagram"
)

// VisualizeForMermaidWithGraphType outputs a visualization of a FSM in Mermaid format as specified by the graphType.
func VisualizeForMermaidWithGraphType(fsm *FSM, graphType MermaidDiagramType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func visualizeForMermaidAsStateDiagram(fsm *FSM) string { _ = "STUB: not implemented"; return "" }

// visualizeForMermaidAsFlowChart outputs a visualization of a FSM in Mermaid format (including highlighting of current state).
func visualizeForMermaidAsFlowChart(fsm *FSM) string { _ = "STUB: not implemented"; return "" }

func writeFlowChartGraphType(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func writeFlowChartStates(buf *bytes.Buffer, sortedStates []string, statesToIDMap map[string]string) {
	_ = "STUB: not implemented"
	return
}

func writeFlowChartTransitions(buf *bytes.Buffer, transitions map[eKey]string, sortedTransitionKeys []eKey, statesToIDMap map[string]string) {
	_ = "STUB: not implemented"
	return
}

func writeFlowChartHighlightCurrent(buf *bytes.Buffer, current string, statesToIDMap map[string]string) {
	_ = "STUB: not implemented"
	return
}

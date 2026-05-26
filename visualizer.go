package fsm

// VisualizeType the type of the visualization
type VisualizeType string

const (
	// GRAPHVIZ the type for graphviz output (http://www.webgraphviz.com/)
	GRAPHVIZ VisualizeType = "graphviz"
	// MERMAID the type for mermaid output (https://mermaid-js.github.io/mermaid/#/stateDiagram) in the stateDiagram form
	MERMAID VisualizeType = "mermaid"
	// MermaidStateDiagram the type for mermaid output (https://mermaid-js.github.io/mermaid/#/stateDiagram) in the stateDiagram form
	MermaidStateDiagram VisualizeType = "mermaid-state-diagram"
	// MermaidFlowChart the type for mermaid output (https://mermaid-js.github.io/mermaid/#/flowchart) in the flow chart form
	MermaidFlowChart VisualizeType = "mermaid-flow-chart"
)

// VisualizeWithType outputs a visualization of a FSM in the desired format.
// If the type is not given it defaults to GRAPHVIZ
func VisualizeWithType(fsm *FSM, visualizeType VisualizeType) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func getSortedTransitionKeys(transitions map[eKey]string) []eKey {
	_ = "STUB: not implemented"
	// we sort the key alphabetically to have a reproducible graph output
	return nil
}

func getSortedStates(transitions map[eKey]string) ([]string, map[string]string) {
	_ = "STUB: not implemented"
	return nil, nil
}

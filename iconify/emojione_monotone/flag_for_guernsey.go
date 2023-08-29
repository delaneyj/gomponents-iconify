package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForGuernsey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2zm27.369 24.1H37.9V4.631C48.611 6.938 57.062 15.389 59.369 26.1zM36.917 54.125h-9.834l2.459-2.458V34.458H12.333l-2.458 2.459v-9.834l2.458 2.459h17.209V12.333l-2.459-2.458h9.834l-2.459 2.458v17.209h17.209l2.458-2.459v9.834l-2.458-2.459H34.458v17.209l2.459 2.458zM26.1 4.631V26.1H4.631C6.938 15.389 15.389 6.938 26.1 4.631zM4.631 37.9H26.1v21.469C15.389 57.062 6.938 48.611 4.631 37.9zM37.9 59.369V37.9h21.469C57.062 48.611 48.611 57.062 37.9 59.369z"/>`),
		g.Group(children),
	)
}
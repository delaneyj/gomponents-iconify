package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecisionTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 12V4h-8v3h-4a2.002 2.002 0 0 0-2 2v6h-6v-3H2v8h8v-3h6v6a2.002 2.002 0 0 0 2 2h4v3h8v-8h-8v3h-4V9h4v3ZM8 18H4v-4h4Zm16 4h4v4h-4Zm0-16h4v4h-4Z"/>`),
		g.Group(children),
	)
}
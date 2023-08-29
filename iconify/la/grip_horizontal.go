package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GripHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 7v8h8V7H2zm10 0v8h8V7h-8zm10 0v8h8V7h-8zM4 9h4v4H4V9zm10 0h4v4h-4V9zm10 0h4v4h-4V9zM2 17v8h8v-8H2zm10 0v8h8v-8h-8zm10 0v8h8v-8h-8zM4 19h4v4H4v-4zm10 0h4v4h-4v-4zm10 0h4v4h-4v-4z"/>`),
		g.Group(children),
	)
}
package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 8v16h22v-3.375l4.563 2.28l1.437.72V8.375l-1.438.72L24 11.374V8H2zm2 2h18v12H4V10zm24 1.625v8.75l-4-2v-4.75l4-2z"/>`),
		g.Group(children),
	)
}
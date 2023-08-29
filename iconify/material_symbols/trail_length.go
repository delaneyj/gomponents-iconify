package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrailLength(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17v-2h6q-.35-.425-.563-.925T12.1 13H9v-2h3.1q.125-.575.338-1.075T13 9H3V7h14q2.075 0 3.538 1.463T22 12q0 2.075-1.463 3.538T17 17H7Zm-5-4v-2h6v2H2Zm1 4v-2h3v2H3Z"/>`),
		g.Group(children),
	)
}
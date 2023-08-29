package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Money(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 4v8H1V4h14zm1-1H0v10h16V3z"/><path fill="currentColor" d="M8 5c1.7 0 3 1.3 3 3s-1.3 3-3 3h5v-1h1V6h-1V5H8zM5 8c0-1.7 1.3-3 3-3H3v1H2v4h1v1h5c-1.7 0-3-1.3-3-3z"/>`),
		g.Group(children),
	)
}
package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 12h9v1H6v-1zm-4.9 1h1.2L6 8L2.3 3H1l3.8 5z"/>`),
		g.Group(children),
	)
}
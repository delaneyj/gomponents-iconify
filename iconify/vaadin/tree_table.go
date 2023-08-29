package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 10V8H4V7h1V5H2v2h1v6h3v-2H4v-1z"/><path fill="currentColor" d="M0 0v16h16V0H0zm7 15H1V3h6v12zm4 0H8V3h3v12zm4 0h-3V3h3v12z"/>`),
		g.Group(children),
	)
}
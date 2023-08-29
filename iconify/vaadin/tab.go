package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 2v12h16V2H0zm13 9h-1V8l-3 3V9H3V7h6V5l3 3V5h1v6z"/>`),
		g.Group(children),
	)
}
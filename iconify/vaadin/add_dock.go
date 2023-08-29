package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddDock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 11v5h16v-5H0zm12 4H9v-3h3v3zm0-8V5c0-5-8-5-8-5s5 0 5 5v2H7l3.5 3L14 7h-2z"/>`),
		g.Group(children),
	)
}
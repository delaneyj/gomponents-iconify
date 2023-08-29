package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HarddriveO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 12h1v1H2v-1zm2 0h3v1H4v-1z"/><path fill="currentColor" d="M13 1H3l-3 9v5h16v-5l-3-9zM3.7 2h8.6l2.7 8H1.1l2.6-8zM1 14v-3h14v3H1z"/>`),
		g.Group(children),
	)
}
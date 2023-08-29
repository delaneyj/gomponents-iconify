package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Download(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 10h-5.5L8 12.5L5.5 10H0v6h16v-6zM4 14H2v-2h2v2z"/><path fill="currentColor" d="M10 6V0H6v6H3l5 5l5-5z"/>`),
		g.Group(children),
	)
}
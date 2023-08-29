package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 10v2H5v-2H0v6h16v-6h-5zm-7 4H2v-2h2v2z"/><path fill="currentColor" d="M13 5L8 0L3 5h3v6h4V5z"/>`),
		g.Group(children),
	)
}
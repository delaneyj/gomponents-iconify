package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 14H0V2h16v12zM1 13h14V3H1v10z"/><path fill="currentColor" d="M2 10v2h12v-1s.2-1.7-2-2c-1.9-.3-2.2.6-3.8.6C7.1 9.6 7.3 8 5 8c-1.7 0-3 2-3 2zm11-4a2 2 0 1 1-3.999.001A2 2 0 0 1 13 6z"/>`),
		g.Group(children),
	)
}
package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExternalBrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 10L8.1 6.8L4.8 10H7v1.8c0 1.7-.9 4.2-4 4.2c4.8 0 6-1.4 6-4.3V10h2z"/><path fill="currentColor" d="M0 0v13h6v-1H1V3h14v9h-5v1h6V0H0zm2 2H1V1h1v1zm11 0H3V1h10v1z"/>`),
		g.Group(children),
	)
}
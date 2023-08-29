package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h1v-1h1v1h12v-1h1v1h1V0H0zm2 14H1v-1h1v1zm0-2H1v-1h1v1zm0-2H1V9h1v1zm0-2H1V7h1v1zm0-2H1V5h1v1zm0-2H1V3h1v1zm0-2H1V1h1v1zm11 13H3V9h10v6zm0-8H3V1h10v6zm2 7h-1v-1h1v1zm0-2h-1v-1h1v1zm0-2h-1V9h1v1zm0-2h-1V7h1v1zm0-2h-1V5h1v1zm0-2h-1V3h1v1zm0-2h-1V1h1v1z"/>`),
		g.Group(children),
	)
}
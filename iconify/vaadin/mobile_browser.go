package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileBrowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 0H3v5H0v11h7v-3h9V0zM6 1h9v1H6V1zM4 1h1v1H4V1zm0 14H3v-1h1v1zm2-2H1V6h5v7zm9-1H7V5H4V3h11v9z"/>`),
		g.Group(children),
	)
}
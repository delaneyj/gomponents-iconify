package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3 5v3h10V5H3zm4 2H4V6h3v1zM3 4h10l-2-4H5zm0 8h10V9H3v3zm8-2h1v1h-1v-1zm-2 0h1v1H9v-1zm-6 6h10v-3H3v3zm1-2h3v1H4v-1z"/>`),
		g.Group(children),
	)
}
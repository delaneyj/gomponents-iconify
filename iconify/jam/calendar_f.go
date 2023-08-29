package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 3v2a2 2 0 1 0 4 0V3h6v2a2 2 0 1 0 4 0V3a3 3 0 0 1 3 3v2H0V6a3 3 0 0 1 3-3zm17 7v6a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3v-6h20zM15 0a1 1 0 0 1 1 1v4a1 1 0 0 1-2 0V1a1 1 0 0 1 1-1zM5 0a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0V1a1 1 0 0 1 1-1z"/>`),
		g.Group(children),
	)
}
package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AltA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M14 7V6h-1V5h-1v1h-.5v1h.5v3.56c0 1 .56 1.44 2 1.44v-1a.899.899 0 0 1-.998-.495L13 7h1zM9 3h1v9H9V3zm-6 9l.57-2h2.82L7 12h1L5.73 4H4.27L2 12h1zm2-6.9L6.11 9H3.89z"/>`),
		g.Group(children),
	)
}
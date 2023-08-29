package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrushAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm6-17H6a1 1 0 0 0-1 1v9a3 3 0 0 0 3 3h1v2.37a4 4 0 1 0 6 0V14h1a3 3 0 0 0 3-3V2a1 1 0 0 0-1-1Zm-6 20a2 2 0 0 1-1.33-3.48a1 1 0 0 0 .33-.74V14h2v2.78a1 1 0 0 0 .33.74A2 2 0 0 1 12 21Zm5-10a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-1h10Zm0-3H7V3h10Z"/>`),
		g.Group(children),
	)
}
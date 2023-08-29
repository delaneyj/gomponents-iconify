package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassNorthEast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M4 6h2v1.5h1V10h1V6h2v10H8v-2H7v-2H6v4H4m8-10h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		g.Group(children),
	)
}
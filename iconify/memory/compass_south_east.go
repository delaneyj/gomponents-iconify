package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassSouthEast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M5 6h5v2H6v2h3v1h1v4H9v1H4v-2h4v-2H5v-1H4V7h1m7-1h6v2h-4v2h4v2h-4v2h4v2h-6"/>`),
		g.Group(children),
	)
}
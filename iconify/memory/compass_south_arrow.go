package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassSouthArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M10 6h2V5h1V4h1V3h1V2h1V0h-2v1h-1v1h-1v1h-2V2H9V1H8V0H6v2h1v1h1v1h1v1h1M9 8h5v2h-4v2h3v1h1v4h-1v1H8v-2h4v-2H9v-1H8V9h1"/>`),
		g.Group(children),
	)
}
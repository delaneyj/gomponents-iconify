package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParkingFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3h7a6 6 0 0 1 0 12h-3v6H6V3Zm4 4v4h3a2 2 0 1 0 0-4h-3Z"/>`),
		g.Group(children),
	)
}
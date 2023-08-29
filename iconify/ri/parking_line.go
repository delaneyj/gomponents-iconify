package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParkingLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3h7a6 6 0 0 1 0 12H8v6H6V3Zm2 2v8h5a4 4 0 0 0 0-8H8Z"/>`),
		g.Group(children),
	)
}
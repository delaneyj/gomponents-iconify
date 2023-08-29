package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSevenThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m171.83 49.15l-48 160A4 4 0 0 1 120 212a3.82 3.82 0 0 1-1.15-.17a4 4 0 0 1-2.68-5L162.62 52H88a4 4 0 0 1 0-8h80a4 4 0 0 1 3.83 5.15Z"/>`),
		g.Group(children),
	)
}
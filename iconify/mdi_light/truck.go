package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 14a2.5 2.5 0 0 1 2.45 2H15V6H4a2 2 0 0 0-2 2v8h1.05a2.5 2.5 0 0 1 2.45-2m0 5a2.5 2.5 0 0 1-2.45-2H1V8a3 3 0 0 1 3-3h11a1 1 0 0 1 1 1v2h3l3 4v5h-2.05a2.5 2.5 0 0 1-4.9 0h-7.1a2.5 2.5 0 0 1-2.45 2m0-4A1.5 1.5 0 0 0 4 16.5A1.5 1.5 0 0 0 5.5 18A1.5 1.5 0 0 0 7 16.5A1.5 1.5 0 0 0 5.5 15m12-1a2.5 2.5 0 0 1 2.45 2H21v-3.68l-.24-.32H16v2.5c.42-.31.94-.5 1.5-.5m0 1a1.5 1.5 0 0 0-1.5 1.5a1.5 1.5 0 0 0 1.5 1.5a1.5 1.5 0 0 0 1.5-1.5a1.5 1.5 0 0 0-1.5-1.5M16 9v2h4l-1.5-2H16Z"/>`),
		g.Group(children),
	)
}
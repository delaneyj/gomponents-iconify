package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingMosqueTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.25 1.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0Zm8.25.75a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm-9 .5a.5.5 0 0 1 .5.5V6h1.17C3.06 5.687 3 5.35 3 5c0-.499.225-.934.497-1.282c.274-.35.627-.65.96-.89A8.345 8.345 0 0 1 5.76 2.06l.024-.012l.007-.003h.002l.002-.001a.5.5 0 0 1 .409 0l.003.001l.007.003l.024.012l.087.041a8.345 8.345 0 0 1 1.217.726c.333.241.686.542.96.891C8.775 4.066 9 4.501 9 5c0 .35-.06.687-.17 1H10V3.5a.5.5 0 0 1 1 0V10a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V9a1 1 0 0 0-2 0v1a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V3.5a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}
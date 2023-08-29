package bxl

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.4 17.8l1.21 2.07L19.77 22L22 14.08L20.72 12L22 10l-2.23-8l-8.16 2.13L10.4 6.2H8L2 12l6 5.81zm9.92-5.8l-1.73 6L15 12l3.59-6zM10.6 6.54L16.84 5l-3.59 6H6.08zM13.27 13l3.59 6l-6.26-1.55L6.1 13z"/>`),
		g.Group(children),
	)
}
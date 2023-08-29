package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShapePolygonPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 15.7V13h2v4l-9 4l-7-7l4-9h4v2H8.3l-2.9 6.6l5 5l6.6-2.9M22 5v2h-3v3h-2V7h-3V5h3V2h2v3h3Z"/>`),
		g.Group(children),
	)
}
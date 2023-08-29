package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BezierCurve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M4 30h8v8H4zm32 0h8v8h-8zM20 10h8v8h-8z"/><path stroke-linecap="round" d="M20 14H4m40 0H28M8 30c0-7.455 5.1-13.72 12-15.496m8 0C34.9 16.28 40 22.544 40 30"/></g>`),
		g.Group(children),
	)
}
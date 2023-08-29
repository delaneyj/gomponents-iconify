package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListBulletedTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15.5L7.5 20h-5L5 15.5M9 19h12v-2H9v2M5 9.5L7.5 14h-5L5 9.5M9 13h12v-2H9v2M5 3.5L7.5 8h-5L5 3.5M9 7h12V5H9v2Z"/>`),
		g.Group(children),
	)
}
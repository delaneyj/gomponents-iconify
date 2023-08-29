package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatListBulletedType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 9.5L7.5 14h-5L5 9.5M3 4h4v4H3V4m2 16a2 2 0 0 0 2-2a2 2 0 0 0-2-2a2 2 0 0 0-2 2a2 2 0 0 0 2 2M9 5v2h12V5H9m0 14h12v-2H9v2m0-6h12v-2H9v2Z"/>`),
		g.Group(children),
	)
}
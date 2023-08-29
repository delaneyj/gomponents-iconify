package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M16 35L10 41L4 35"/><path d="M16 13L10 7L4 13"/><path d="M10 7V41"/><path d="M44 9H22"/><path d="M36 19H22"/><path d="M44 29H22"/><path d="M36 39H22"/></g>`),
		g.Group(children),
	)
}
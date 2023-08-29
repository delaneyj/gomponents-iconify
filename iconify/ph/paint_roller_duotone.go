package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaintRollerDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 64v64a8 8 0 0 1-8 8H48a8 8 0 0 1-8-8V64a8 8 0 0 1 8-8h152a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M232 88h-16V64a16 16 0 0 0-16-16H48a16 16 0 0 0-16 16v24H16a8 8 0 0 0 0 16h16v24a16 16 0 0 0 16 16h152a16 16 0 0 0 16-16v-24h16v50l-100.4 28.65A16.07 16.07 0 0 0 120 198v34a8 8 0 0 0 16 0v-34l100.4-28.68A16.07 16.07 0 0 0 248 154v-50a16 16 0 0 0-16-16Zm-32 40H48V64h152v64Z"/></g>`),
		g.Group(children),
	)
}
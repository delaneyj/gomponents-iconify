package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChafingDishOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M6 18H42V24C42 27.3137 39.3137 30 36 30H12C8.68629 30 6 27.3137 6 24V18Z"/><path d="M40 42H8"/><path d="M13 42L16 30"/><path d="M35 42L32 30"/><path d="M30 18L27 4H21L18 18"/><path d="M36 11H40"/><path d="M8 11H12"/></g>`),
		g.Group(children),
	)
}
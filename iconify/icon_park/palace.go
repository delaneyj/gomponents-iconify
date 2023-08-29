package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Palace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M4 18H44L24 6L4 18Z"/><path d="M44 42L4 42"/><path d="M9 18V42"/><path d="M19 18V42"/><path d="M29 18V42"/><path d="M39 18V42"/></g>`),
		g.Group(children),
	)
}
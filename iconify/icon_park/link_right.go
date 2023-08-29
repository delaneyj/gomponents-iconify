package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 12V4H24V12"/><path d="M24 36V44H8V36"/><path d="M44 24L24 24"/><path d="M16 34V14"/><path d="M36 16L44 24L36 32"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cactus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 43h32M6 16c0 5.5.5 14 10 14m26-16c0 5 0 20-10 20"/><path d="M24 5a8 8 0 0 0-8 8v30h16V13a8 8 0 0 0-8-8Z"/></g>`),
		g.Group(children),
	)
}
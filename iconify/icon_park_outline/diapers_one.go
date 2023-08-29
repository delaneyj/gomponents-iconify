package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiapersOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 6H6v18c0 4.5 2.5 18.5 18 18.5S42 28 42 24V6ZM6 14h10m16 0h10"/><path d="M42 24c-10 0-17 4.8-17 18M6 24c10 0 17 4.8 17 18"/></g>`),
		g.Group(children),
	)
}
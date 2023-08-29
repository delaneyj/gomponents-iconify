package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelEightTitle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M6 8v32M24 8v32M7 24h16"/><path d="M37 29a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 11a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g>`),
		g.Group(children),
	)
}
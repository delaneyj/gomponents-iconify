package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TitleLevel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 8v32M24 8v32M7 24h16"/><path d="M32 24v16m0-8.976C32 28.46 34 26 37 26s5 2.358 5 5.024v8.99"/></g>`),
		g.Group(children),
	)
}
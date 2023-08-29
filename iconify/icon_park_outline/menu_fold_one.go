package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuFoldOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M8 10.5h32m-16 9h16m-16 9h16m-32 9h32"/><path d="m8 19l8 5l-8 5V19Z"/></g>`),
		g.Group(children),
	)
}
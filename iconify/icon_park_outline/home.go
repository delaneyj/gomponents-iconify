package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M9 18v24h30V18L24 6L9 18Z"/><path stroke-linejoin="round" d="M19 29v13h10V29H19Z"/><path stroke-linecap="round" d="M9 42h30"/></g>`),
		g.Group(children),
	)
}
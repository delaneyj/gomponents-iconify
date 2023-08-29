package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M44 10H4v30h40V10Z"/><path stroke-linecap="round" d="m10 16l28 18m0-18L24 35m0-19L10 34"/></g>`),
		g.Group(children),
	)
}
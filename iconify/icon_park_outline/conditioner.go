package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conditioner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 24h32v20H8zm29 0v-7H11v7m6-7c2-2.167 3-6.5 3-13c3 0 10 5.417 10 12.733"/><path d="M16 31h16v6H16z"/></g>`),
		g.Group(children),
	)
}
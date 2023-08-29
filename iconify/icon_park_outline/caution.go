package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Caution(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M24 5L2 43h44L24 5Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M24 35v1m0-17l.008 10"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListMiddle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M8 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0-16a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 28a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-linecap="round" d="M20 24h24M20 38h24M20 10h24"/></g>`),
		g.Group(children),
	)
}
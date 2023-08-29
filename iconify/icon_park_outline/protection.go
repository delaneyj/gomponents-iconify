package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Protection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M42 24c0-9.389-8.059-17-18-17S6 14.611 6 24h36Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M24 24.008v14.537C24 41.558 21.514 44 18.5 44S13 41.558 13 38.545"/><path stroke-linecap="round" d="M24 4v3"/></g>`),
		g.Group(children),
	)
}
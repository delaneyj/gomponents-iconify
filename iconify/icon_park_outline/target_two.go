package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M24 41c9.39 0 17-7.61 17-17S33.39 7 24 7S7 14.61 7 24s7.61 17 17 17Z"/><path d="M24 32c4.42 0 8-3.58 8-8s-3.58-8-8-8s-8 3.58-8 8s3.58 8 8 8Z"/><path stroke-linecap="round" d="M4 24h40M24 4v40"/></g>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridSixteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><rect width="40" height="40" x="4" y="4" stroke-linejoin="round" rx="2"/><path d="M14 4v40M24 4v40M34 4v40"/><path stroke-linejoin="round" d="M4 14h40M4 34h40M4 24h40"/></g>`),
		g.Group(children),
	)
}
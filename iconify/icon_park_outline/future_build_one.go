package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FutureBuildOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M14 24c1.217-8.296 10-19 10-19s8.782 10.704 10 19c1.09 7.432-3 20-3 20H17s-4.091-12.568-3-20Z"/><path d="M18 14h12m-15 6h18m-19 6h20m-19 6h18m-17 6h16"/><path stroke-linejoin="round" d="M4 44h40"/><path d="M24 4v2"/></g>`),
		g.Group(children),
	)
}
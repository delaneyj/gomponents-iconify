package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StartTimeSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5v25h36V5"/><path stroke-linejoin="round" d="M30.037 10L25 15.014l5.037 5.098M30 37l-6 6l-6-6m6-7v13"/><path d="M20 9.002V21"/></g>`),
		g.Group(children),
	)
}
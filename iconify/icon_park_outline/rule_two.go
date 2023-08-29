package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M24 13C11 13 4 22.85 4 35h40c0-12.15-7-22-20-22Z"/><path stroke-linecap="round" d="M10 31v4m7-4v4m7-4v4m7-4v4m7-4v4"/><path d="M24 20c-4.554 0-10 2.134-10 6h20c0-3.866-5.446-6-10-6Z"/></g>`),
		g.Group(children),
	)
}
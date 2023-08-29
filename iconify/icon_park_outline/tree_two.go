package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M11 24s-3 2-3 6c0 6 10 8 10 8m19-14s3 2 3 6c0 6-10 8-10 8m-6-22.001v28"/><path d="M19.126 13.5a6 6 0 1 1 9.748 0"/><path stroke-linejoin="round" d="M13 13s-4 2.5-4 6s3 6 3 6m23-12s4 2.5 4 6s-3 6-3 6"/></g>`),
		g.Group(children),
	)
}
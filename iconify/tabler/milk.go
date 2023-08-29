package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Milk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 6h8V4a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v2zm8 0l1.094 1.759a6 6 0 0 1 .906 3.17V19a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-8.071a6 6 0 0 1 .906-3.17L8 6"/><path d="M10 16a2 2 0 1 0 4 0a2 2 0 1 0-4 0m0-6h4"/></g>`),
		g.Group(children),
	)
}
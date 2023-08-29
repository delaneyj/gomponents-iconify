package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpZeroOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 8l4-4l4 4M7 4v16"/><rect width="4" height="6" x="15" y="4" ry="2"/><path d="M17 20v-6h-2m0 6h4"/></g>`),
		g.Group(children),
	)
}
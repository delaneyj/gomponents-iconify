package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownOneZero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 16l4 4l4-4m-4 4V4m10 6V4h-2m0 6h4"/><rect width="4" height="6" x="15" y="14" ry="2"/></g>`),
		g.Group(children),
	)
}
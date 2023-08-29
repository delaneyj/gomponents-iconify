package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M11.25 20a.75.75 0 0 0 1.5 0h-1.5Zm1.5 0V4h-1.5v16h1.5Z" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18 10l-6-6l-6 6"/></g>`),
		g.Group(children),
	)
}
package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftUpLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17.47 18.53a.75.75 0 1 0 1.06-1.06l-1.06 1.06Zm1.06-1.06l-12-12l-1.06 1.06l12 12l1.06-1.06Z" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 6H6v9"/></g>`),
		g.Group(children),
	)
}
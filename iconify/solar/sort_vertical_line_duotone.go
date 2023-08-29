package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortVerticalLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 18V6m0 0l4 4.125M16 6l-4 4.125" opacity=".5"/><path d="M8 6v12m0 0l4-4.125M8 18l-4-4.125"/></g>`),
		g.Group(children),
	)
}
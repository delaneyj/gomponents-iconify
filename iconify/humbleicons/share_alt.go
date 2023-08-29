package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.437 7.912a2.5 2.5 0 1 0 4.127-2.824a2.5 2.5 0 0 0-4.127 2.824Zm0 0l-4.96 3.394m0 0a3 3 0 1 0 .236 2.979m-.237-2.98c.33.483.524 1.066.524 1.695c0 .46-.103.895-.288 1.285m0 0l4.528 2.145m0 0a2.5 2.5 0 1 0 4.52 2.141a2.5 2.5 0 0 0-4.52-2.142Z"/>`),
		g.Group(children),
	)
}
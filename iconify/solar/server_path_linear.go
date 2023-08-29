package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServerPathLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M22 19h-8M2 19h8m2-2v-3"/><circle cx="12" cy="19" r="2" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-width="1.5" d="M2 11a3 3 0 0 1 3-3h14a3 3 0 1 1 0 6H5a3 3 0 0 1-3-3Zm0-6a3 3 0 0 1 3-3h14a3 3 0 1 1 0 6H5a3 3 0 0 1-3-3Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M13 5h6m-6 6h6"/><circle cx="6" cy="5" r="1" fill="currentColor"/><circle cx="6" cy="11" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}
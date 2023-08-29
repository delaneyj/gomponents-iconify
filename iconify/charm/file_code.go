package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 7.75v-6h5.5l5 5v7.5h-2"/><path d="M7.75 2.25v5h5.05M6.75 10.8l2 1.75l-2 1.75m-3-3.5l-2 1.75l2 1.75"/></g>`),
		g.Group(children),
	)
}
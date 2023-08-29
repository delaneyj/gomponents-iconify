package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCompare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12.5" cy="12.5" r="1.75"/><circle cx="3.5" cy="3.5" r="1.75"/><path d="M3.75 5.75v5c0 1 .5 1.5 1.5 1.5h2m-.5 2l1.5-2l-1.5-2m5.5 0v-5c0-1-.5-1.5-1.5-1.5h-2m.5-2l-1.5 2l1.5 2"/></g>`),
		g.Group(children),
	)
}
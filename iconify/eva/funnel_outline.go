package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaFunnelOutline0"><g id="evaFunnelOutline1"><path id="evaFunnelOutline2" fill="currentColor" d="M13.9 22a1 1 0 0 1-.6-.2l-4-3.05a1 1 0 0 1-.39-.8v-3.27l-4.8-9.22A1 1 0 0 1 5 4h14a1 1 0 0 1 .86.49a1 1 0 0 1 0 1l-5 9.21V21a1 1 0 0 1-.55.9a1 1 0 0 1-.41.1Zm-3-4.54l2 1.53v-4.55A1 1 0 0 1 13 14l4.3-8H6.64l4.13 8a1 1 0 0 1 .11.46Z"/></g></g>`),
		g.Group(children),
	)
}
package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaFunnelFill0"><g id="evaFunnelFill1"><path id="evaFunnelFill2" fill="currentColor" d="M13.9 22a1 1 0 0 1-.6-.2l-4-3.05a1 1 0 0 1-.39-.8v-3.27l-4.8-9.22A1 1 0 0 1 5 4h14a1 1 0 0 1 .86.49a1 1 0 0 1 0 1l-5 9.21V21a1 1 0 0 1-.55.9a1 1 0 0 1-.41.1Z"/></g></g>`),
		g.Group(children),
	)
}
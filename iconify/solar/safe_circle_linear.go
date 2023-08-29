package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafeCircleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Z"/><path stroke-linecap="round" d="M7 8v8"/><path d="M12 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z"/><path stroke-linecap="round" d="m16 10l1-1m-6 6l1-1m0-4l-1-1m6 6l-1-1"/></g>`),
		g.Group(children),
	)
}
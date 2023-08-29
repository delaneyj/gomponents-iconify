package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoPointsCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm14 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="m19 19l-1.5-1.5m-2-2l-1-1m-2-2l-1-1m-2-2l-1-1m-2-2L5 5"/></g>`),
		g.Group(children),
	)
}
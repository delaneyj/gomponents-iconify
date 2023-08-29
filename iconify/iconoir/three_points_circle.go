package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreePointsCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M5 10.5V9m0 6v-1.5"/><path fill="currentColor" d="M5 20a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm14 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M10.5 19H9m6 0h-1.5"/></g>`),
		g.Group(children),
	)
}
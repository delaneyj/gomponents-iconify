package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlingBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path fill="currentColor" d="M11.5 8a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm-4 3a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Zm4 2a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1Z"/></g>`),
		g.Group(children),
	)
}
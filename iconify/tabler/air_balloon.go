package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirBalloon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 20a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1zm2-4c3.314 0 6-4.686 6-8A6 6 0 1 0 6 8c0 3.314 2.686 8 6 8z"/><path d="M10 9a2 7 0 1 0 4 0a2 7 0 1 0-4 0"/></g>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 10.111V1l11 6v14H3V7l7 3.111Zm2-5.742v8.82l-7-3.111V19h14V8.187L12 4.37Z"/>`),
		g.Group(children),
	)
}
package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterDropF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 .565c4.667 6.09 7 10.423 7 13a7 7 0 1 1-14 0c0-2.577 2.333-6.91 7-13z"/>`),
		g.Group(children),
	)
}
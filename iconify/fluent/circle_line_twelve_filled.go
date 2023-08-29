package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleLineTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 6a5 5 0 0 1 10 0H1Zm.1 1a5.002 5.002 0 0 0 9.8 0H1.1Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassNorthwestTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 26c6.627 0 12-5.373 12-12S20.627 2 14 2S2 7.373 2 14s5.373 12 12 12Zm-.424-8.702a5 5 0 0 1-2.874-2.874L8.783 9.43a.5.5 0 0 1 .646-.647l4.995 1.92a5 5 0 0 1 2.874 2.873l1.919 4.995a.5.5 0 0 1-.646.646l-4.995-1.919Z"/>`),
		g.Group(children),
	)
}
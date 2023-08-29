package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFourPointsSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.74 10.75L12 8l1.25 2.75L16 12l-2.75 1.26L12 16l-1.26-2.74L8 12l2.74-1.25Z"/>`),
		g.Group(children),
	)
}
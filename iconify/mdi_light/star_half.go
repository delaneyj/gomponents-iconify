package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 6.06l-1.86 4.39l-4.75.41L8 14l-1.08 4.63L11 16.18v1.16L5.42 20.7l1.46-6.35l-4.92-4.28l6.49-.57l2.55-6v2.56Z"/>`),
		g.Group(children),
	)
}
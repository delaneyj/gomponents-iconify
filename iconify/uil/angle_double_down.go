package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.29 11.46a1 1 0 0 0 1.42 0l3-3A1 1 0 1 0 14.29 7L12 9.34L9.71 7a1 1 0 1 0-1.42 1.46Zm3 1.08L12 14.84l-2.29-2.3A1 1 0 0 0 8.29 14l3 3a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0-1.42-1.42Z"/>`),
		g.Group(children),
	)
}
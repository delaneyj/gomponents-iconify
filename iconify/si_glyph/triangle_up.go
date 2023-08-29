package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.96 2.392a1.49 1.49 0 0 1 2.104 0l6.442 6.444c.582.581.839 2.103-1 2.103H2.518c-1.902 0-1.582-1.521-1.001-2.103L7.96 2.392z"/>`),
		g.Group(children),
	)
}
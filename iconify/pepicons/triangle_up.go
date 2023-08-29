package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.848 13.47l-5-8a1 1 0 0 0-1.696 0l-5 8A1 1 0 0 0 5 15h10a1 1 0 0 0 .848-1.53ZM6.804 13L10 7.887L13.196 13H6.804Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
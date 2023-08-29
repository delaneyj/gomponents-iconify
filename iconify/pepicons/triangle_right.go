package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.53 15.848l8-5a1 1 0 0 0 0-1.696l-8-5A1 1 0 0 0 6 5v10a1 1 0 0 0 1.53.848ZM8 13.196V6.804L13.113 10L8 13.196Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
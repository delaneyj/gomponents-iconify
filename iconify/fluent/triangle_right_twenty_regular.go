package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.215 11.318c1.047-.567 1.047-2.07 0-2.638L5.216 2.18A1.5 1.5 0 0 0 3.002 3.5v12.999a1.5 1.5 0 0 0 2.214 1.319l11.999-6.5Zm-.476-1.758a.5.5 0 0 1 0 .879l-11.999 6.5a.5.5 0 0 1-.738-.44v-13a.5.5 0 0 1 .738-.439l11.999 6.5Z"/>`),
		g.Group(children),
	)
}
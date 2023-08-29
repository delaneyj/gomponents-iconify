package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiningTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 8l7.5 4l-7.5 4l-4 7.5L8 16L.5 12L8 8l4-7.5L16 8Zm3.25 4l-4.728-2.522L12 4.75L9.478 9.478L4.75 12l4.728 2.522L12 19.25l2.522-4.728L19.25 12Z"/>`),
		g.Group(children),
	)
}
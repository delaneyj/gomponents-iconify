package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.682 17.215c.567 1.047 2.07 1.047 2.637 0l6.5-12A1.5 1.5 0 0 0 16.5 3.003H3.501a1.5 1.5 0 0 0-1.319 2.214l6.5 11.998Zm1.758-.477a.5.5 0 0 1-.879 0L3.061 4.74a.5.5 0 0 1 .44-.74H16.5a.5.5 0 0 1 .44.739l-6.5 11.998Z"/>`),
		g.Group(children),
	)
}
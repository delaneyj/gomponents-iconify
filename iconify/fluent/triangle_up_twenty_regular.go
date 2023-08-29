package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleUpTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.682 2.785c.567-1.047 2.07-1.047 2.637 0l6.5 12a1.5 1.5 0 0 1-1.319 2.213H3.501a1.5 1.5 0 0 1-1.319-2.214l6.5-11.999Zm1.758.477a.5.5 0 0 0-.879 0l-6.5 11.998a.5.5 0 0 0 .44.739H16.5a.5.5 0 0 0 .44-.739l-6.5-11.998Z"/>`),
		g.Group(children),
	)
}
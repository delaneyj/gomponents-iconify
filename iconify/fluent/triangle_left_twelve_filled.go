package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.459 5.214a.903.903 0 0 0 0 1.572l7.169 4.092c.61.348 1.372-.089 1.372-.787V1.91c0-.698-.762-1.135-1.372-.787l-7.17 4.092Z"/>`),
		g.Group(children),
	)
}
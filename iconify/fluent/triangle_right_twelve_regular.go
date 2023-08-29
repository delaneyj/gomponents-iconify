package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.541 5.214a.903.903 0 0 1 0 1.572l-7.169 4.092C2.762 11.226 2 10.789 2 10.09V1.91c0-.698.762-1.135 1.372-.787l7.17 4.092ZM9.881 6L3.017 2.083v7.834L9.881 6Z"/>`),
		g.Group(children),
	)
}
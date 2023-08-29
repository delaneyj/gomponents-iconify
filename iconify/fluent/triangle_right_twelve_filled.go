package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.541 6.786a.903.903 0 0 0 0-1.572L3.372 1.122C2.762.774 2 1.211 2 1.91v8.18c0 .698.762 1.135 1.372.787l7.17-4.092Z"/>`),
		g.Group(children),
	)
}
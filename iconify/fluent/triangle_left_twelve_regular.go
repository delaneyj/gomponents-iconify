package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleLeftTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1.459 6.786a.903.903 0 0 1 0-1.572l7.169-4.092C9.238.774 10 1.211 10 1.91v8.182c0 .698-.762 1.135-1.372.787l-7.17-4.092ZM2.119 6l6.864 3.917V2.083L2.119 6Z"/>`),
		g.Group(children),
	)
}
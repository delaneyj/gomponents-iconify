package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.312 14.223a1.5 1.5 0 0 1-2.629 0l-5.5-10a1.5 1.5 0 0 1 1.315-2.222h10.999a1.5 1.5 0 0 1 1.314 2.223l-5.5 9.999Zm-1.753-.482a.5.5 0 0 0 .877 0l5.499-10a.5.5 0 0 0-.438-.74H2.498a.5.5 0 0 0-.438.74l5.5 10Z"/>`),
		g.Group(children),
	)
}
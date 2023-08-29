package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveUpRightTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.146 6.854a.5.5 0 0 0 .708-.708l-4-4a.5.5 0 0 0-.708 0l-4 4a.5.5 0 1 0 .708.708L10 3.707V10c0 1.965-.247 3.38-.764 4.473c-.512 1.08-1.308 1.887-2.493 2.598a.5.5 0 0 0 .514.858c1.315-.79 2.269-1.732 2.882-3.027c.608-1.283.861-2.867.861-4.902V3.707l3.146 3.147Z"/>`),
		g.Group(children),
	)
}
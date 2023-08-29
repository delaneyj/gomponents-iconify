package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveUpLeftTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.854 6.854a.5.5 0 1 1-.708-.708l4-4a.5.5 0 0 1 .708 0l4 4a.5.5 0 0 1-.708.708L10 3.707V10c0 1.965.247 3.38.764 4.473c.512 1.08 1.308 1.887 2.493 2.598a.5.5 0 0 1-.514.858c-1.315-.79-2.269-1.732-2.882-3.027C9.253 13.619 9 12.035 9 10V3.707L5.854 6.854Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleOffSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m11.874 12.582l2.272 2.272a.5.5 0 0 0 .708-.708l-13-13a.5.5 0 1 0-.708.708l2.272 2.272a6 6 0 0 0 8.456 8.456ZM14 8a5.972 5.972 0 0 1-.83 3.048L4.951 2.83A6 6 0 0 1 14 8Z"/>`),
		g.Group(children),
	)
}
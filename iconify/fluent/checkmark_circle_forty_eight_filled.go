package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckmarkCircleFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 4c11.046 0 20 8.954 20 20s-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4Zm8.634 13.616a1.25 1.25 0 0 0-1.666-.091l-.102.091L20.75 27.732l-3.616-3.616a1.25 1.25 0 0 0-1.859 1.666l.091.102l4.5 4.5a1.25 1.25 0 0 0 1.666.091l.102-.091l11-11a1.25 1.25 0 0 0 0-1.768Z"/>`),
		g.Group(children),
	)
}
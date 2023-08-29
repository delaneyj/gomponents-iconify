package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowULeftBottomBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10.5c0 4.14-3.36 7.5-7.5 7.5H11v4l-7-6l7-6v4h2.5c1.93 0 3.5-1.57 3.5-3.5S15.43 7 13.5 7H6V3h7.5c4.14 0 7.5 3.36 7.5 7.5Z"/>`),
		g.Group(children),
	)
}
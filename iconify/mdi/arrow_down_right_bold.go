package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 3c4.14 0 7.5 3.36 7.5 7.5V13h4l-6 7l-6-7h4v-2.5C14 8.57 12.43 7 10.5 7H3V3h7.5Z"/>`),
		g.Group(children),
	)
}
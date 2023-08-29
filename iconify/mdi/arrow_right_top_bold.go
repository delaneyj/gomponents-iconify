package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightTopBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13.5C3 9.36 6.36 6 10.5 6H13V2l7 6l-7 6v-4h-2.5C8.57 10 7 11.57 7 13.5V21H3v-7.5Z"/>`),
		g.Group(children),
	)
}
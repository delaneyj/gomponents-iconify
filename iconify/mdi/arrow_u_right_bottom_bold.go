package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightBottomBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 3H18v4h-7.5C8.57 7 7 8.57 7 10.5S8.57 14 10.5 14H13v-4l7 6l-7 6v-4h-2.5C6.36 18 3 14.64 3 10.5S6.36 3 10.5 3Z"/>`),
		g.Group(children),
	)
}
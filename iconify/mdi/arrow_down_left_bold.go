package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3v4h-7.5C11.57 7 10 8.57 10 10.5V13h4l-6 7l-6-7h4v-2.5C6 6.36 9.36 3 13.5 3H21Z"/>`),
		g.Group(children),
	)
}
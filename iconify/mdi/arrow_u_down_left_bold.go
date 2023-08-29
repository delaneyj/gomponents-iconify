package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUDownLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 10.5V18h-4v-7.5C17 8.57 15.43 7 13.5 7S10 8.57 10 10.5V13h4l-6 7l-6-7h4v-2.5C6 6.36 9.36 3 13.5 3S21 6.36 21 10.5Z"/>`),
		g.Group(children),
	)
}
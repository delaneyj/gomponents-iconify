package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarFifteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M14 7a1.5 1.5 0 0 0-1.15-1.45l-1.39-3.24A.5.5 0 0 0 11 2H4a.5.5 0 0 0-.44.28L2.15 5.54A1.5 1.5 0 0 0 1 7v3.5h1v1a1 1 0 1 0 2 0v-1h7v1a1 1 0 1 0 2 0v-1h1V7zM4.3 3h6.4l1.05 2.5h-8.5L4.3 3zM3 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm9 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}
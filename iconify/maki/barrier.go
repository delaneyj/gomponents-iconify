package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barrier(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13 2H2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h1v2.5a.5.5 0 0 0 1 0V12h7v.5a.5.5 0 0 0 1 0V10h1a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zm0 1v2l-2-2zM9.5 3L13 6.5V9L7 3zm-4 6L2 5.5V3l6 6zM2 9V7l2 2zm9 2H4v-1h7zm-.207-2H9.5l-6-6h2l6 6z"/>`),
		g.Group(children),
	)
}
package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DifferenceSet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 40V19a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2ZM42 8v21a2 2 0 0 1-2 2h-7a2 2 0 0 1-2-2V19a2 2 0 0 0-2-2H19a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h21a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}
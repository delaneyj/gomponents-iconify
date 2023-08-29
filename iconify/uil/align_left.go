package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Zm0 4h14a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Zm18 2H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2Zm-4 4H3a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
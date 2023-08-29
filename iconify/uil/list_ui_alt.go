package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUiAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 2h14a1 1 0 0 0 0-2h-14a1 1 0 0 0 0 2Zm0 3a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm4 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm10-5h-10a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2Zm0 5h-6a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
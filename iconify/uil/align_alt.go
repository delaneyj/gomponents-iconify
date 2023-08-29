package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 11H3a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2Zm0 8H3a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2Zm4-14h7a1 1 0 0 0 0-2h-7a1 1 0 0 0 0 2Zm-4 2H3a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2Zm0 8H3a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2Zm0-12H7a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2Zm11 4h-7a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2Zm0 4h-7a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2Zm-4 8h-3a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2Zm4-4h-7a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 5H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm0 4H3a1 1 0 0 1 0-2h18a1 1 0 0 1 0 2zm-4 4H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2z"/>`),
		g.Group(children),
	)
}
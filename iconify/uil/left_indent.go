package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftIndent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Zm0 4h10a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Zm18.77-1.31a1 1 0 0 0-1.41-.12l-2 1.66a1 1 0 0 0 0 1.54l2 1.66a1 1 0 0 0 .64.24a1 1 0 0 0 .77-.36a1 1 0 0 0-.13-1.41l-1.08-.9l1.08-.9a1 1 0 0 0 .13-1.41ZM21 17H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2ZM3 15h10a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2Z"/>`),
		g.Group(children),
	)
}
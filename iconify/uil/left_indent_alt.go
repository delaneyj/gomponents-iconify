package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftIndentAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 5a1 1 0 0 0-1 1v12a1 1 0 0 0 2 0V6a1 1 0 0 0-1-1Zm4 2h8a1 1 0 0 0 0-2h-8a1 1 0 0 0 0 2ZM5.77 9.69a1 1 0 0 0-1.41-.13l-2 1.67a1 1 0 0 0 0 1.54l2 1.67a1 1 0 0 0 1.41-.13a1 1 0 0 0-.13-1.41L4.56 12l1.08-.9a1 1 0 0 0 .13-1.41ZM21 9h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2Zm0 4h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2Zm0 4h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
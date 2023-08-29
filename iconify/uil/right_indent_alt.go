package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightIndentAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.64 9.56a1 1 0 1 0-1.28 1.54l1.08.9l-1.08.9a1 1 0 0 0-.13 1.41a1 1 0 0 0 1.41.13l2-1.67a1 1 0 0 0 0-1.54ZM9 5a1 1 0 0 0-1 1v12a1 1 0 0 0 2 0V6a1 1 0 0 0-1-1Zm4 2h8a1 1 0 0 0 0-2h-8a1 1 0 0 0 0 2Zm8 10h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2Zm0-8h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2Zm0 4h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}
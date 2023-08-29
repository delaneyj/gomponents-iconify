package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5a1 1 0 0 0 0 2h16a1 1 0 1 0 0-2H4Zm0 8a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2H4Zm2-3a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1Zm1 7a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7Z"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5a1 1 0 1 1 0 2H4a1 1 0 0 1 0-2h16Zm0 4a1 1 0 1 1 0 2h-8a1 1 0 1 1 0-2h8Zm1 5a1 1 0 0 0-1-1H4a1 1 0 1 0 0 2h16a1 1 0 0 0 1-1Zm-1 3a1 1 0 1 1 0 2h-8a1 1 0 1 1 0-2h8Z"/>`),
		g.Group(children),
	)
}
package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5a1 1 0 0 0 0 2h16a1 1 0 1 0 0-2H4Zm0 4a1 1 0 0 0 0 2h16a1 1 0 1 0 0-2H4Zm-1 5a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Zm1 3a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2H4Z"/>`),
		g.Group(children),
	)
}
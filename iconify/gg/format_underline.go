package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatUnderline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 10V4h2v6a4 4 0 0 0 8 0V4h2v6a6 6 0 0 1-12 0Zm1 8a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2H7Z"/>`),
		g.Group(children),
	)
}
package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHorizontalAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 8v3H3v2h10v3l4-4l-4-4m6 12h2V4h-2v16Z"/>`),
		g.Group(children),
	)
}
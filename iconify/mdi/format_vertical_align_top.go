package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatVerticalAlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11h3v10h2V11h3l-4-4l-4 4M4 3v2h16V3H4Z"/>`),
		g.Group(children),
	)
}
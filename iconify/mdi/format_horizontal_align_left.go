package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHorizontalAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 16v-3h10v-2H11V8l-4 4l4 4m-8 4h2V4H3v16Z"/>`),
		g.Group(children),
	)
}
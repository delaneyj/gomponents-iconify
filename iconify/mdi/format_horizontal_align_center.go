package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHorizontalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 16v-3h4v-2h-4V8l-4 4l4 4M5 8v3H1v2h4v3l4-4l-4-4m6 12h2V4h-2v16Z"/>`),
		g.Group(children),
	)
}
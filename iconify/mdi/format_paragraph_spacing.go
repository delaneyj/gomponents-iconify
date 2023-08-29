package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatParagraphSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17h18v2H3v-2M3 2h18v2H3V2m0 18h18v2H3v-2M13 8h2l-3-3l-3 3h2v5H9l3 3l3-3h-2V8Z"/>`),
		g.Group(children),
	)
}
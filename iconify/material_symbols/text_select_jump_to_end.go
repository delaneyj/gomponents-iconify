package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectJumpToEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 16l-1.4-1.4l1.575-1.6H3v-2h7.175L8.6 9.4L10 8l4 4l-4 4Zm9 5V3h2v18h-2ZM15 5V3h2v2h-2Zm0 16v-2h2v2h-2ZM11 5V3h2v2h-2Zm0 16v-2h2v2h-2ZM7 5V3h2v2H7Zm0 16v-2h2v2H7ZM3 5V3h2v2H3Zm0 16v-2h2v2H3Z"/>`),
		g.Group(children),
	)
}
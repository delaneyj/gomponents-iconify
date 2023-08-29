package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoSizeSelectLargeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V7h14v14H3Zm2-3h10l-3.4-4.5L9 17l-1.6-2.15L5 18ZM3 5V3h2v2H3Zm8 0V3h2v2h-2Zm8 0V3h2v2h-2ZM7 5V3h2v2H7Zm12 8v-2h2v2h-2Zm0 8v-2h2v2h-2Zm0-12V7h2v2h-2Zm0 8v-2h2v2h-2ZM15 5V3h2v2h-2Z"/>`),
		g.Group(children),
	)
}
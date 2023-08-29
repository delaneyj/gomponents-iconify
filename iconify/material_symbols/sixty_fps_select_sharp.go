package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixtyFpsSelectSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 14V4h7v2H6v2h5v6H4Zm2-2h3v-2H6v2Zm7 2V4h7v10h-7Zm2-2h3V6h-3v6ZM3 22v-5h2v5H3Zm4 0v-5h2v5H7Zm4 0v-5h2v5h-2Zm4 0v-5h6v5h-6Z"/>`),
		g.Group(children),
	)
}
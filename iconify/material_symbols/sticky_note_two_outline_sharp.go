package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteTwoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h9v-5h5V5H5v14Zm-2 2V3h18v12l-6 6H3Zm4-7v-2h5v2H7Zm0-4V8h10v2H7Zm-2 9V5v14Z"/>`),
		g.Group(children),
	)
}
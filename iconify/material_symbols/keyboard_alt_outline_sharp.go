package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardAltOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V4h22v17H1Zm2-2h18V6H3v13Zm5-2h8v-1H8v1Zm-3-3h2v-2H5v2Zm4 0h2v-2H9v2Zm4 0h2v-2h-2v2Zm4 0h2v-2h-2v2ZM5 10h2V8H5v2Zm4 0h2V8H9v2Zm4 0h2V8h-2v2Zm4 0h2V8h-2v2ZM3 19V6v13Z"/>`),
		g.Group(children),
	)
}
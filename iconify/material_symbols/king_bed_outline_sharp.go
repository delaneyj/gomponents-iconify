package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KingBedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19H4l-.65-2H2v-7h2V5h16v5h2v7h-1.35L20 19h-1l-.65-2H5.65L5 19Zm8-9h5V7h-5v3Zm-7 0h5V7H6v3Zm-2 5h16v-3H4v3Zm16 0H4h16Z"/>`),
		g.Group(children),
	)
}
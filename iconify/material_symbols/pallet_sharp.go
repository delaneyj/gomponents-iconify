package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PalletSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-4h20v4h-3v-2h-5.5v2h-3v-2H5v2H2Zm3-6V2h14v14H5Zm4-8h6V6H9v2Z"/>`),
		g.Group(children),
	)
}
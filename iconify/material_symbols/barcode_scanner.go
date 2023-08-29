package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarcodeScanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21v-5h2v3h3v2H1Zm17 0v-2h3v-3h2v5h-5ZM4 18V6h2v12H4Zm3 0V6h1v12H7Zm3 0V6h2v12h-2Zm3 0V6h3v12h-3Zm4 0V6h1v12h-1Zm2 0V6h1v12h-1ZM1 8V3h5v2H3v3H1Zm20 0V5h-3V3h5v5h-2Z"/>`),
		g.Group(children),
	)
}
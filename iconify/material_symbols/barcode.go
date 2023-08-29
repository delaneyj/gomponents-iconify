package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h2v14H1Zm3 0V5h2v14H4Zm3 0V5h1v14H7Zm3 0V5h2v14h-2Zm3 0V5h3v14h-3Zm4 0V5h1v14h-1Zm3 0V5h3v14h-3Z"/>`),
		g.Group(children),
	)
}
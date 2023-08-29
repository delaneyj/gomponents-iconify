package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentScannerOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 6V1h5v2H4v3H2Zm18 0V3h-3V1h5v5h-2ZM2 23v-5h2v3h3v2H2Zm15 0v-2h3v-3h2v5h-5ZM7 18h10V6H7v12Zm-2 2V4h14v16H5Zm4-10h6V8H9v2Zm0 3h6v-2H9v2Zm0 3h6v-2H9v2Zm-2 2V6v12Z"/>`),
		g.Group(children),
	)
}
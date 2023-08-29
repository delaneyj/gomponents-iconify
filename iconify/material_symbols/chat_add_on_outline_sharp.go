package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatAddOnOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20V3h16v7.075q-.25-.05-.5-.063T18 10q-.25 0-.5.013t-.5.062V5H5v10h7.075q-.05.25-.063.5T12 16q0 .25.013.5t.062.5H6l-3 3ZM7 9h8V7H7v2Zm0 4h5v-2H7v2Zm10 7v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM5 15V5v10Z"/>`),
		g.Group(children),
	)
}
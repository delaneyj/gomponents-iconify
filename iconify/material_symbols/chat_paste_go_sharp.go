package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatPasteGoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 21l-1.4-1.4l1.575-1.6H14v-2h4.175L16.6 14.4L18 13l4 4l-4 4ZM7 10h8V8H7v2Zm0 4h5v-2H7v2Zm-4 7V4h16v7.075q-.25-.05-.5-.063T18 11q-2.525 0-4.263 1.75T12 17q0 .25.013.5t.062.5H6l-3 3Z"/>`),
		g.Group(children),
	)
}
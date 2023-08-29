package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentPasteGoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 21l-1.4-1.425L18.175 18H12v-2h6.175L16.6 14.4L18 13l4 4l-4 4Zm3-11h-2V5h-2v3H7V5H5v14h5v2H3V3h6.175q.275-.875 1.075-1.438T12 1q1 0 1.788.563T14.85 3H21v7Zm-9-5q.425 0 .713-.288T13 4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4q0 .425.288.713T12 5Z"/>`),
		g.Group(children),
	)
}
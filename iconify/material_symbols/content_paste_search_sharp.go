package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentPasteSearchSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.6 23l-2.7-2.7q-.55.325-1.15.513T16.5 21q-1.875 0-3.187-1.313T12 16.5q0-1.875 1.313-3.188T16.5 12q1.875 0 3.188 1.313T21 16.5q0 .65-.188 1.25T20.3 18.9l2.7 2.7l-1.4 1.4Zm-5.1-4q1.05 0 1.775-.725T19 16.5q0-1.05-.725-1.775T16.5 14q-1.05 0-1.775.725T14 16.5q0 1.05.725 1.775T16.5 19Zm4.5-9h-2V5h-2v3H7V5H5v14h5v2H3V3h6.175q.275-.875 1.075-1.438T12 1q1 0 1.788.563T14.85 3H21v7Zm-9-5q.425 0 .713-.288T13 4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4q0 .425.288.713T12 5Z"/>`),
		g.Group(children),
	)
}
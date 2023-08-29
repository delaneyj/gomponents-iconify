package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NfcOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h10V7h-4q-.825 0-1.413.588T11 9v1.3q-.5.275-.75.7T10 12q0 .825.588 1.413T12 14q.825 0 1.413-.588T14 12q0-.575-.275-1T13 10.3V9h2v6H9V9h1V7H7v10Zm-4 4V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}
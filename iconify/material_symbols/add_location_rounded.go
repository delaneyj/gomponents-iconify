package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddLocationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 11v2q0 .425.288.713T12 14q.425 0 .713-.288T13 13v-2h2q.425 0 .713-.288T16 10q0-.425-.288-.713T15 9h-2V7q0-.425-.288-.713T12 6q-.425 0-.713.288T11 7v2H9q-.425 0-.713.288T8 10q0 .425.288.713T9 11h2Zm1 10.325q-.35 0-.7-.125t-.625-.375Q9.05 19.325 7.8 17.9t-2.087-2.762q-.838-1.338-1.275-2.575T4 10.2q0-3.75 2.413-5.975T12 2q3.175 0 5.588 2.225T20 10.2q0 1.125-.438 2.363t-1.275 2.575Q17.45 16.475 16.2 17.9t-2.875 2.925q-.275.25-.625.375t-.7.125Z"/>`),
		g.Group(children),
	)
}
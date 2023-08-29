package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 15H2q-.425 0-.713-.288T1 14q0-.425.288-.713T2 13h20q.425 0 .713.288T23 14q0 .425-.288.713T22 15Zm-4-6l-5-5v4q0 .425.288.713T14 9h4ZM6 22q-.825 0-1.413-.588T4 20v-3h16v3q0 .825-.588 1.413T18 22H6ZM4 11V4q0-.825.588-1.413T6 2h7.175q.4 0 .763.15t.637.425l4.85 4.85q.275.275.425.638t.15.762V11H4Z"/>`),
		g.Group(children),
	)
}
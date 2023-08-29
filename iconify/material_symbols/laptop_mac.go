package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopMac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20q-.825 0-1.413-.588T0 18h4q-.825 0-1.413-.588T2 16V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v11q0 .825-.588 1.413T20 18h4q0 .825-.588 1.413T22 20H2Zm10-1q.425 0 .713-.288T13 18q0-.425-.288-.713T12 17q-.425 0-.713.288T11 18q0 .425.288.713T12 19Z"/>`),
		g.Group(children),
	)
}
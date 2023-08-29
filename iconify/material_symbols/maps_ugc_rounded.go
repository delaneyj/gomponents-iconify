package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapsUgcRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.95 16.3q-.475-1.025-.713-2.1T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22q-1.125 0-2.2-.238t-2.1-.712L2.75 22.5q-.575.175-1-.25t-.25-1l1.45-4.95ZM11 13v2q0 .425.288.713T12 16q.425 0 .713-.288T13 15v-2h2q.425 0 .713-.288T16 12q0-.425-.288-.713T15 11h-2V9q0-.425-.288-.713T12 8q-.425 0-.713.288T11 9v2H9q-.425 0-.713.288T8 12q0 .425.288.713T9 13h2Z"/>`),
		g.Group(children),
	)
}
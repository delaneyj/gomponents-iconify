package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatterySaverRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 19h-2q-.425 0-.713-.288T13 18q0-.425.288-.713T14 17h2v-2q0-.425.288-.713T17 14q.425 0 .713.288T18 15v2h2q.425 0 .713.288T21 18q0 .425-.288.713T20 19h-2v2q0 .425-.288.713T17 22q-.425 0-.713-.288T16 21v-2Zm-8 3q-.425 0-.713-.288T7 21V5q0-.425.288-.713T8 4h2V3q0-.425.288-.713T11 2h2q.425 0 .713.288T14 3v1h2q.425 0 .713.288T17 5v7q-2.5.025-4.25 1.763T11 18q0 1.15.4 2.175T12.525 22H8Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18q-.425 0-.713-.288T7 17q0-.425.288-.713T8 16q.425 0 .713.288T9 17q0 .425-.288.713T8 18Zm4 0q-.425 0-.713-.288T11 17q0-.425.288-.713T12 16q.425 0 .713.288T13 17q0 .425-.288.713T12 18Zm4 0q-.425 0-.713-.288T15 17q0-.425.288-.713T16 16q.425 0 .713.288T17 17q0 .425-.288.713T16 18ZM5 14v-2q0-2.65 1.7-4.6T11 5.1V3h2v2.1q2.6.35 4.3 2.3T19 12v2H5Zm2-2h10q0-2.075-1.463-3.538T12 7Q9.925 7 8.462 8.463T7 12Zm1 9q-.425 0-.713-.288T7 20q0-.425.288-.713T8 19q.425 0 .713.288T9 20q0 .425-.288.713T8 21Zm4 0q-.425 0-.713-.288T11 20q0-.425.288-.713T12 19q.425 0 .713.288T13 20q0 .425-.288.713T12 21Zm4 0q-.425 0-.713-.288T15 20q0-.425.288-.713T16 19q.425 0 .713.288T17 20q0 .425-.288.713T16 21Zm-4-9Z"/>`),
		g.Group(children),
	)
}
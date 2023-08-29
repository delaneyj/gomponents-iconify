package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FamiliarFaceAndZoneOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-3.35 0-5.675-2.325T4 12q0-3.35 2.325-5.675T12 4q3.35 0 5.675 2.325T20 12q0 3.35-2.325 5.675T12 20Zm0-2q2.5 0 4.25-1.75T18 12q0-.425-.063-.838t-.187-.812q-.375.075-.75.113t-.75.037q-1.575 0-3-.6T10.7 8.15q-.7 1.425-1.925 2.475T6 12.15q.075 2.45 1.812 4.15T12 18ZM6.4 9.85q1.1-.575 1.675-1.337T9.2 6.7q-.95.5-1.675 1.313T6.4 9.85ZM9.5 14q-.425 0-.713-.287T8.5 13q0-.425.288-.713T9.5 12q.425 0 .713.288T10.5 13q0 .425-.288.713T9.5 14Zm6.75-5.5h.3q.15 0 .3-.025q-.825-1.125-2.087-1.8T12 6h-.3q-.15 0-.275.025q.975 1.125 2.062 1.8t2.763.675ZM14.5 14q-.425 0-.713-.287T13.5 13q0-.425.288-.713T14.5 12q.425 0 .713.288T15.5 13q0 .425-.288.713T14.5 14ZM1 6V1h5v2H3v3H1Zm0 17v-5h2v3h3v2H1Zm17 0v-2h3v-3h2v5h-5Zm3-17V3h-3V1h5v5h-2Zm-9.575.025ZM9.2 6.7Z"/>`),
		g.Group(children),
	)
}
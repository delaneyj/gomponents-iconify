package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FamiliarFaceAndZone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q-3.35 0-5.675-2.325T4 12q0-3.35 2.325-5.675T12 4q3.35 0 5.675 2.325T20 12q0 3.35-2.325 5.675T12 20Zm0-2q2.5 0 4.25-1.75T18 12q0-.425-.063-.838t-.187-.812q-.375.075-.75.113t-.75.037q-1.575 0-3-.6T10.7 8.15q-.7 1.425-1.925 2.475T6 12.15q.075 2.45 1.812 4.15T12 18Zm-2.5-4q-.425 0-.713-.288T8.5 13q0-.425.288-.713T9.5 12q.425 0 .713.288T10.5 13q0 .425-.288.713T9.5 14Zm5 0q-.425 0-.713-.288T13.5 13q0-.425.288-.713T14.5 12q.425 0 .713.288T15.5 13q0 .425-.288.713T14.5 14ZM1 6V3q0-.825.588-1.413T3 1h3v2H3v3H1Zm5 17H3q-.825 0-1.413-.588T1 21v-3h2v3h3v2Zm12 0v-2h3v-3h2v3q0 .825-.588 1.413T21 23h-3Zm3-17V3h-3V1h3q.825 0 1.413.588T23 3v3h-2Z"/>`),
		g.Group(children),
	)
}
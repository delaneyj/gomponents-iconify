package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PriceChangeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm6-6H7q-.425 0-.713.288T6 15q0 .425.288.713T7 16h1q0 .425.288.713T9 17q.425 0 .713-.288T10 16h1q.425 0 .713-.288T12 15v-3q0-.425-.288-.713T11 11H8v-1h3q.425 0 .713-.288T12 9q0-.425-.288-.713T11 8h-1q0-.425-.288-.713T9 7q-.4 0-.563.363T8 8H7q-.425 0-.713.288T6 9v3q0 .425.288.713T7 13h3v1Zm6.175 2.075l1.4-1.4q.125-.125.063-.275t-.238-.15h-2.8q-.175 0-.237.15t.062.275l1.4 1.4q.075.075.175.075t.175-.075ZM14.6 10h2.8q.175 0 .238-.15t-.063-.275l-1.4-1.4Q16.1 8.1 16 8.1t-.175.075l-1.4 1.4q-.125.125-.062.275t.237.15Z"/>`),
		g.Group(children),
	)
}
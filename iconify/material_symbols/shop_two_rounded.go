package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopTwoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.825 0-1.413-.588T1 20V10q0-.425.288-.713T2 9q.425 0 .713.288T3 10v10h16q.425 0 .713.288T20 21q0 .425-.288.713T19 22H3Zm4-4q-.825 0-1.413-.588T5 16V6q0-.425.288-.713T6 5h4V3q0-.825.588-1.413T12 1h4q.825 0 1.413.588T18 3v2h4q.425 0 .713.288T23 6v10q0 .825-.588 1.413T21 18H7Zm5-13h4V3h-4v2Zm0 9.225q0 .3.263.425t.512-.025l4.05-2.575q.225-.15.225-.425t-.225-.425l-4.05-2.575q-.25-.15-.513-.025T12 9.025v5.2Z"/>`),
		g.Group(children),
	)
}
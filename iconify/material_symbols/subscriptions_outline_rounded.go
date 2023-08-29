package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptionsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22q-.825 0-1.413-.588T2 20V10q0-.825.588-1.413T4 8h16q.825 0 1.413.588T22 10v10q0 .825-.588 1.413T20 22H4Zm0-2h16V10H4v10Zm6.775-1.525l4.6-3.05q.225-.15.225-.425t-.225-.425l-4.6-3.05q-.25-.175-.513-.038t-.262.438v6.15q0 .3.263.438t.512-.038ZM5 7q-.425 0-.713-.288T4 6q0-.425.288-.713T5 5h14q.425 0 .713.288T20 6q0 .425-.288.713T19 7H5Zm3-3q-.425 0-.713-.288T7 3q0-.425.288-.713T8 2h8q.425 0 .713.288T17 3q0 .425-.288.713T16 4H8ZM4 20V10v10Z"/>`),
		g.Group(children),
	)
}
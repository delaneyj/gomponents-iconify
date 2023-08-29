package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPostOfficeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 19v2q0 .425-.288.713T15 22H3q-.425 0-.713-.288T2 21v-8q0-.425.288-.713T3 12h3V8q0-2.5 1.75-4.25T12 2h4q2.5 0 4.25 1.75T22 8v14h-2v-3h-4Zm0-2h4V8q0-1.65-1.175-2.825T16 4h-4q-1.65 0-2.825 1.175T8 8v4h7q.425 0 .713.288T16 13v4Zm-6-7V8h8v2h-8Zm-1 6.85L14 14H4l5 2.85Zm0 1.75l-5-2.85V20h10v-4.25L9 18.6ZM4 14v6v-6Z"/>`),
		g.Group(children),
	)
}
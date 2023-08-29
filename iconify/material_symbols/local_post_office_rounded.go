package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalPostOfficeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21v-8q0-.425.288-.713T3 12h3V8q0-2.5 1.75-4.25T12 2h4q2.5 0 4.25 1.75T22 8v13q0 .425-.288.713T21 22q-.425 0-.713-.288T20 21v-2h-4v2q0 .425-.288.713T15 22H3Zm6-5.15l-3.875-2.2q-.375-.225-.75 0T4 15.3q0 .2.1.388t.275.287l4.125 2.35q.225.125.5.125t.5-.125l4.125-2.35q.175-.1.275-.287t.1-.388q0-.425-.375-.65t-.75 0L9 16.85Zm7 .15h4V8q0-1.65-1.175-2.825T16 4h-4q-1.65 0-2.825 1.175T8 8v4h7q.425 0 .713.288T16 13v4Zm-5-7q-.425 0-.713-.288T10 9q0-.425.288-.713T11 8h6q.425 0 .713.288T18 9q0 .425-.288.713T17 10h-6Z"/>`),
		g.Group(children),
	)
}
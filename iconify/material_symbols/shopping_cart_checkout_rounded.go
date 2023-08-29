package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoppingCartCheckoutRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22q-.825 0-1.413-.588T5 20q0-.825.588-1.413T7 18q.825 0 1.413.588T9 20q0 .825-.588 1.413T7 22Zm10 0q-.825 0-1.413-.588T15 20q0-.825.588-1.413T17 18q.825 0 1.413.588T19 20q0 .825-.588 1.413T17 22ZM3 4H2q-.425 0-.713-.288T1 3q0-.425.288-.713T2 2h1.65q.275 0 .525.15t.375.425L8.525 11h7l3.625-6.5q.125-.25.35-.375T20 4q.575 0 .863.488t.012.987L17.3 11.95q-.275.5-.738.775T15.55 13H8.1L7 15h11q.425 0 .713.288T19 16q0 .425-.288.713T18 17H7q-1.125 0-1.713-.975T5.25 14.05L6.6 11.6L3 4Zm9.175 3H9q-.425 0-.713-.288T8 6q0-.425.288-.713T9 5h3.175l-.9-.9q-.3-.3-.288-.7t.313-.7q.3-.275.7-.288t.7.288l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.687.288T11.3 9.3q-.275-.275-.275-.7t.275-.7l.875-.9Z"/>`),
		g.Group(children),
	)
}
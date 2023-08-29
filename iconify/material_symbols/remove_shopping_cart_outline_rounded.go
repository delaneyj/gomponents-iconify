package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveShoppingCartOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.825 13l-2-2h1.725l2.75-5H8.825l-2-2H19.95q.575 0 .888.488t.012 1.062l-3.55 6.4q-.275.5-.713.775t-.762.275ZM7 22q-.825 0-1.413-.588T5 20q0-.825.588-1.413T7 18q.825 0 1.413.588T9 20q0 .825-.588 1.413T7 22Zm12.8.6L14.15 17H7.6q-1.1 0-1.675-.938T5.85 14.1l1.05-2.15L5.1 7.9L1.4 4.2q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275ZM12.15 15l-2-2H8.6l-1 2h4.55Zm3.4-4h-1.725h1.725ZM17 22q-.825 0-1.413-.588T15 20q0-.825.588-1.413T17 18q.825 0 1.413.588T19 20q0 .825-.588 1.413T17 22Z"/>`),
		g.Group(children),
	)
}
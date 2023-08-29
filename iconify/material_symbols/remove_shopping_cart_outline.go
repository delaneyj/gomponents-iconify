package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveShoppingCartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.825 13l-2-2h1.725l2.75-5H8.825l-2-2H19.95q.575 0 .888.488t.012 1.062l-3.55 6.4q-.275.5-.713.775t-.762.275ZM7 22q-.825 0-1.413-.588T5 20q0-.825.588-1.413T7 18q.825 0 1.413.588T9 20q0 .825-.588 1.413T7 22Zm13.5 1.3L14.15 17H7.6q-1.1 0-1.675-.938T5.85 14.1l1.05-2.15L5.1 7.9L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM12.15 15l-2-2H8.6l-1 2h4.55Zm3.4-4h-1.725h1.725ZM17 22q-.825 0-1.413-.588T15 20q0-.825.588-1.413T17 18q.825 0 1.413.588T19 20q0 .825-.588 1.413T17 22Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CribRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20q.5 0 1-.063t1-.187V16h-4v3.75q.5.125 1 .188T12 20Zm0 2q-1.725 0-3.325-.575t-2.95-1.65q-.35-.275-.362-.7T5.7 18.3q.275-.275.663-.275t.687.25q.225.2.463.35t.487.3V16H6q-.825 0-1.413-.588T4 14V8q0-1.65 1.175-2.825T8 4h2q.825 0 1.413.588T12 6v3h6q.825 0 1.413.588T20 11v3q0 .825-.588 1.413T18 16h-2v2.925q.25-.15.488-.3t.462-.35q.3-.25.7-.238t.675.288q.325.325.313.75t-.363.7q-1.35 1.075-2.95 1.65T12 22Z"/>`),
		g.Group(children),
	)
}
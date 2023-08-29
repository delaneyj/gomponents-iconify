package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberDvr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15h3.5q.65 0 1.075-.425T9 13.5v-3q0-.65-.425-1.075T7.5 9H4v6Zm1.5-1.5v-3h2v3h-2Zm5.6 1.5h1.5l1.75-6h-1.5l-1 3.45l-1-3.45h-1.5l1.75 6Zm3.9 0h1.5v-2h1.15l.85 2H20l-.9-2.1q.375-.2.638-.575T20 11.5v-1q0-.65-.425-1.075T18.5 9H15v6Zm1.5-3.5v-1h2v1h-2ZM3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3Z"/>`),
		g.Group(children),
	)
}
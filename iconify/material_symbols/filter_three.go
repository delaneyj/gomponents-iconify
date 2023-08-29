package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h4q.825 0 1.413-.588T17 13v-1.5q0-.65-.425-1.075T15.5 10q.65 0 1.075-.425T17 8.5V7q0-.825-.588-1.413T15 5h-4v2h4v2h-2v2h2v2h-4v2Zm-3 3q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V6h2v14h14v2H4Z"/>`),
		g.Group(children),
	)
}
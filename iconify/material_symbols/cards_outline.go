package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 11.5V6h5.5v5.5H6ZM6 18v-5.5h5.5V18H6Zm6.5-6.5V6H18v5.5h-5.5Zm0 6.5v-5.5H18V18h-5.5ZM8 9.5h1.5V8H8v1.5Zm6.5 0H16V8h-1.5v1.5ZM8 16h1.5v-1.5H8V16Zm6.5 0H16v-1.5h-1.5V16Zm-5-6.5Zm5 0Zm0 5Zm-5 0ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14Z"/>`),
		g.Group(children),
	)
}
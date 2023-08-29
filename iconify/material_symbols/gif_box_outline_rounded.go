package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GifBoxOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Zm3.5 9h1q.425 0 .713-.288T10.5 13v-.5q0-.2-.15-.35T10 12q-.2 0-.35.15t-.15.35v.5h-1v-2H10q.2 0 .35-.15t.15-.35q0-.2-.15-.35T10 10H8.5q-.425 0-.713.288T7.5 11v2q0 .425.288.713T8.5 14Zm3.5 0q.2 0 .35-.15t.15-.35v-3q0-.2-.15-.35T12 10q-.2 0-.35.15t-.15.35v3q0 .2.15.35T12 14Zm2 0q.2 0 .35-.15t.15-.35v-1h1q.2 0 .35-.15T16 12q0-.2-.15-.35t-.35-.15h-1V11H16q.2 0 .35-.15t.15-.35q0-.2-.15-.35T16 10h-2q-.2 0-.35.15t-.15.35v3q0 .2.15.35T14 14Z"/>`),
		g.Group(children),
	)
}
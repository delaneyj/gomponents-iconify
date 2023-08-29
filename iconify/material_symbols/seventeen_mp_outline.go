package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeventeenMpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 11.5H10v-6H7V7h1.5v4.5Zm4.75 0H15l1.45-4.7q.15-.5-.15-.9t-.8-.4H12V7h2.6l-1.35 4.5ZM6 18.5h1.5V14h1v3H10v-3h1v4.5h1.5v-5q0-.425-.288-.713T11.5 12.5H7q-.425 0-.713.288T6 13.5v5Zm7.5 0H15V17h2q.425 0 .713-.288T18 16v-2.5q0-.425-.288-.713T17 12.5h-3.5v6Zm1.5-3V14h1.5v1.5H15ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 15H11v-1.5H8v-1h2q.425 0 .713-.288T11 11.5V10q0-.425-.288-.713T10 9H6.5v1.5h3v1h-2q-.425 0-.713.288T6.5 12.5V15Zm6.5 0h4q.425 0 .713-.288T18 14v-4q0-.425-.288-.713T17 9h-4v6Zm1.5-1.5v-3h2v3h-2ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}
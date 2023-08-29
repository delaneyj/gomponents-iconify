package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EighteenUpRatingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 15H10V9H7v1.5h1.5V15Zm4 0H15q.425 0 .713-.288T16 14v-4q0-.425-.288-.713T15 9h-2.5q-.425 0-.713.288T11.5 10v4q0 .425.288.713T12.5 15Zm.5-1v-1.5h1.5V14H13Zm0-2.5V10h1.5v1.5H13ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
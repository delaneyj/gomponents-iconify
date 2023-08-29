package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentPasteSearchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 21q-1.875 0-3.188-1.313T12 16.5q0-1.875 1.313-3.188T16.5 12q1.875 0 3.188 1.313T21 16.5q0 .65-.188 1.25T20.3 18.9l2 2q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2-2q-.55.325-1.15.513T16.5 21Zm0-2q1.05 0 1.775-.725T19 16.5q0-1.05-.725-1.775T16.5 14q-1.05 0-1.775.725T14 16.5q0 1.05.725 1.775T16.5 19ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h4.175q.275-.875 1.075-1.438T12 1q1 0 1.788.563T14.85 3H19q.825 0 1.413.588T21 5v4q0 .425-.288.713T20 10q-.425 0-.713-.288T19 9V5h-2v2q0 .425-.288.713T16 8H8q-.425 0-.713-.288T7 7V5H5v14h4q.425 0 .713.288T10 20q0 .425-.288.713T9 21H5Zm7-16q.425 0 .713-.288T13 4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4q0 .425.288.713T12 5Z"/>`),
		g.Group(children),
	)
}
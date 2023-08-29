package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicExternalOnOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.8 7q-.375-.425-.587-.925T4 5q0-1.25.875-2.125T7 2q1.25 0 2.125.875T10 5q0 .575-.213 1.075T9.2 7H4.8ZM10 22q-1.65 0-2.825-1.175T6 18H5L4 8h6L9 18H8q0 .825.588 1.413T10 20q.825 0 1.413-.588T12 18V6q0-1.65 1.175-2.825T16 2q1.65 0 2.825 1.175T20 6v16h-2V6q0-.825-.588-1.413T16 4q-.825 0-1.413.588T14 6v12q0 1.65-1.175 2.825T10 22Zm-3.2-6h.4l.6-6H6.2l.6 6Zm.4-6h-1h1.6h-.6Z"/>`),
		g.Group(children),
	)
}
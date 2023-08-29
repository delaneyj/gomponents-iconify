package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneIphoneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-5v3h10v-3H7Zm5 2.5q.425 0 .713-.288T13 19.5q0-.425-.288-.713T12 18.5q-.425 0-.713.288T11 19.5q0 .425.288.713T12 20.5ZM7 16h10V6H7v10ZM7 4h10V3H7v1Zm0 14v3v-3ZM7 4V3v1Z"/>`),
		g.Group(children),
	)
}
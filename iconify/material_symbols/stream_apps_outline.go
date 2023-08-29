package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreamAppsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13h5v-3h-5v3Zm-2 4V9q0-.425.288-.713T15 8h7q.425 0 .713.288T23 9v5q0 .425-.288.713T22 15h-6l-2 2Zm-7 6q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v4h-2V6H7v12h10v-1h2v4q0 .825-.588 1.413T17 23H7Zm0-3v1h10v-1H7ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Zm9-7v-3v3Z"/>`),
		g.Group(children),
	)
}
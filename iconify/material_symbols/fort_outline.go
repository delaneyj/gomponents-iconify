package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FortOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21v-4l2-2V9L1 7V3h2v2h2V3h2v2h2V3h2v4L9 9v1h6V9l-2-2V3h2v2h2V3h2v2h2V3h2v4l-2 2v6l2 2v4h-9v-3q0-.825-.588-1.413T12 16q-.825 0-1.413.588T10 18v3H1Zm2-2h5v-1q0-1.65 1.175-2.825T12 14q1.65 0 2.825 1.175T16 18v1h5v-1.175l-2-2v-7.65L20.175 7h-4.35L17 8.175V12H7V8.175L8.175 7h-4.35L5 8.175v7.65l-2 2V19Zm9-6Z"/>`),
		g.Group(children),
	)
}
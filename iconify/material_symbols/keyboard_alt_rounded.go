package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v13q0 .825-.588 1.413T21 21H3Zm5.5-4h7q.2 0 .35-.15t.15-.35q0-.2-.15-.35T15.5 16h-7q-.2 0-.35.15T8 16.5q0 .2.15.35t.35.15ZM5 14h2v-2H5v2Zm4 0h2v-2H9v2Zm4 0h2v-2h-2v2Zm4 0h2v-2h-2v2ZM5 10h2V8H5v2Zm4 0h2V8H9v2Zm4 0h2V8h-2v2Zm4 0h2V8h-2v2Z"/>`),
		g.Group(children),
	)
}
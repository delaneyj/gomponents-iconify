package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 4l2 4h3L7 4h2l2 4h3l-2-4h2l2 4h3l-2-4h3q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4Zm0 6v8h16v-8H4Zm0 0v8v-8Z"/>`),
		g.Group(children),
	)
}
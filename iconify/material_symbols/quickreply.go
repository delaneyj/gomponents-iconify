package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quickreply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 23v-5h-2v-6h5l-1.7 4h2.2L19 23ZM2 22V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v6h-7v8H6l-4 4Z"/>`),
		g.Group(children),
	)
}
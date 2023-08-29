package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V8H4v10Zm3.5-1l-1.4-1.4L8.675 13l-2.6-2.6L7.5 9l4 4l-4 4Zm4.5 0v-2h6v2h-6Z"/>`),
		g.Group(children),
	)
}
package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonsAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V8q0-.825.588-1.413T4 6h16q.825 0 1.413.588T22 8v8q0 .825-.588 1.413T20 18H4Zm0-2h16V8H4v8Zm3.25-1h1.5v-2.25H11v-1.5H8.75V9h-1.5v2.25H5v1.5h2.25V15ZM4 16V8v8Z"/>`),
		g.Group(children),
	)
}
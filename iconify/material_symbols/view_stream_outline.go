package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewStreamOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 17v-4H5v4h14Zm0-6V7H5v4h14ZM5 19q-.825 0-1.413-.588T3 17V7q0-.825.588-1.413T5 5h14q.825 0 1.413.588T21 7v10q0 .825-.588 1.413T19 19H5Z"/>`),
		g.Group(children),
	)
}
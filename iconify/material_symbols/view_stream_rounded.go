package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewStreamRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19q-.825 0-1.413-.588T3 17v-2q0-.825.588-1.413T5 13h14q.825 0 1.413.588T21 15v2q0 .825-.588 1.413T19 19H5Zm0-8q-.825 0-1.413-.588T3 9V7q0-.825.588-1.413T5 5h14q.825 0 1.413.588T21 7v2q0 .825-.588 1.413T19 11H5Z"/>`),
		g.Group(children),
	)
}
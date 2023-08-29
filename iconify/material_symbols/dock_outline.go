package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 23v-2h8v2H8Zm0-4q-.825 0-1.413-.588T6 17V3q0-.825.588-1.413T8 1h8q.825 0 1.413.588T18 3v14q0 .825-.588 1.413T16 19H8Zm0-3v1h8v-1H8Zm0-2h8V6H8v8ZM8 4h8V3H8v1Zm0 0V3v1Zm0 12v1v-1Z"/>`),
		g.Group(children),
	)
}
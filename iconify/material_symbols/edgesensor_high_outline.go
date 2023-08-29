package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgesensorHighOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 17v-7h2v7H0Zm3-3V7h2v7H3Zm5 8q-.825 0-1.413-.588T6 20V4q0-.825.588-1.413T8 2h8q.825 0 1.413.588T18 4v16q0 .825-.588 1.413T16 22H8Zm8-3H8v1h8v-1Zm-8-2h8V7H8v10ZM8 5h8V4H8v1Zm11 12v-7h2v7h-2Zm3-3V7h2v7h-2ZM8 5V4v1Zm0 14v1v-1Z"/>`),
		g.Group(children),
	)
}
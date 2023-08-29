package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 17.5l-4-4v1.675l-2-2V6H8.825l-2-2H16q.825 0 1.413.588T18 6v4.5l4-4v11Zm-1.45 5.85L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4Zm-8.1-13.725ZM9.55 12.4ZM4 4l2 2H4v12h12v-2l2 2q0 .825-.588 1.413T16 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4Z"/>`),
		g.Group(children),
	)
}
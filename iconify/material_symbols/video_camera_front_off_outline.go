package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraFrontOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 17.5l-4-4v1.675l-2-2V6H8.825l-2-2H16q.825 0 1.413.588T18 6v4.5l4-4v11Zm-1.45 5.85L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4ZM12.425 9.575ZM9.6 12.4ZM4 4l2 2H4v12h12v-2l2 2q0 .825-.588 1.413T16 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4Zm2 12v-.55q0-1.1 1.1-1.775T10 13q1.8 0 2.9.675T14 15.45V16H6Z"/>`),
		g.Group(children),
	)
}
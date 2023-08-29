package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BallotOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 10h5V8h-5v2Zm0 6h5v-2h-5v2Zm-3-5q.825 0 1.413-.588T11 9q0-.825-.588-1.413T9 7q-.825 0-1.413.588T7 9q0 .825.588 1.413T9 11Zm0 6q.825 0 1.413-.588T11 15q0-.825-.588-1.413T9 13q-.825 0-1.413.588T7 15q0 .825.588 1.413T9 17Zm-4 4q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}
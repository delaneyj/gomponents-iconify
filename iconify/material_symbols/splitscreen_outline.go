package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitscreenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 9h12V4H6v5Zm0 2q-.825 0-1.413-.588T4 9V4q0-.825.588-1.413T6 2h12q.825 0 1.413.588T20 4v5q0 .825-.588 1.413T18 11H6Zm0 9h12v-5H6v5Zm0 2q-.825 0-1.413-.588T4 20v-5q0-.825.588-1.413T6 13h12q.825 0 1.413.588T20 15v5q0 .825-.588 1.413T18 22H6ZM6 9V4v5Zm0 11v-5v5Z"/>`),
		g.Group(children),
	)
}
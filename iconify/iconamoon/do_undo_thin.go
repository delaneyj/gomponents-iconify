package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoUndoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.354 6.354a.5.5 0 1 0-.708-.708l.708.708ZM5 9l-.354-.354a.5.5 0 0 0 0 .708L5 9Zm2.646 3.354a.5.5 0 0 0 .708-.708l-.708.708ZM5 16.5a.5.5 0 0 0 0 1v-1ZM7.646 5.646l-3 3l.708.708l3-3l-.708-.708Zm-3 3.708l3 3l.708-.708l-3-3l-.708.708ZM5 9.5h11v-1H5v1ZM19.5 13a3.5 3.5 0 0 1-3.5 3.5v1a4.5 4.5 0 0 0 4.5-4.5h-1ZM16 9.5a3.5 3.5 0 0 1 3.5 3.5h1A4.5 4.5 0 0 0 16 8.5v1Zm-11 8h11v-1H5v1Z"/>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoRedoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.354 5.646a.5.5 0 0 0-.708.708l.708-.708ZM19 9l.354.354a.5.5 0 0 0 0-.708L19 9Zm-3.354 2.646a.5.5 0 0 0 .708.708l-.708-.708ZM19 17.5a.5.5 0 0 0 0-1v1ZM15.646 6.354l3 3l.708-.708l-3-3l-.708.708Zm3 2.292l-3 3l.708.708l3-3l-.708-.708ZM19 8.5H8v1h11v-1ZM3.5 13A4.5 4.5 0 0 0 8 17.5v-1A3.5 3.5 0 0 1 4.5 13h-1ZM8 8.5A4.5 4.5 0 0 0 3.5 13h1A3.5 3.5 0 0 1 8 9.5v-1Zm11 8H8v1h11v-1Z"/>`),
		g.Group(children),
	)
}
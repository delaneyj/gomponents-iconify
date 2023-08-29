package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoUndoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.53 6.53a.75.75 0 0 0-1.06-1.06l1.06 1.06ZM5 9l-.53-.53a.75.75 0 0 0 0 1.06L5 9Zm2.47 3.53a.75.75 0 0 0 1.06-1.06l-1.06 1.06ZM5 16.25a.75.75 0 0 0 0 1.5v-1.5ZM7.47 5.47l-3 3l1.06 1.06l3-3l-1.06-1.06Zm-3 4.06l3 3l1.06-1.06l-3-3l-1.06 1.06Zm.53.22h11v-1.5H5v1.5ZM19.25 13A3.25 3.25 0 0 1 16 16.25v1.5A4.75 4.75 0 0 0 20.75 13h-1.5ZM16 9.75A3.25 3.25 0 0 1 19.25 13h1.5A4.75 4.75 0 0 0 16 8.25v1.5Zm-11 8h11v-1.5H5v1.5Z"/>`),
		g.Group(children),
	)
}
package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoRedoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.53 5.47a.75.75 0 1 0-1.06 1.06l1.06-1.06ZM19 9l.53.53a.75.75 0 0 0 0-1.06L19 9Zm-3.53 2.47a.75.75 0 1 0 1.06 1.06l-1.06-1.06ZM19 17.75a.75.75 0 0 0 0-1.5v1.5ZM15.47 6.53l3 3l1.06-1.06l-3-3l-1.06 1.06Zm3 1.94l-3 3l1.06 1.06l3-3l-1.06-1.06Zm.53-.22H8v1.5h11v-1.5ZM3.25 13A4.75 4.75 0 0 0 8 17.75v-1.5A3.25 3.25 0 0 1 4.75 13h-1.5ZM8 8.25A4.75 4.75 0 0 0 3.25 13h1.5A3.25 3.25 0 0 1 8 9.75v-1.5Zm11 8H8v1.5h11v-1.5Z"/>`),
		g.Group(children),
	)
}
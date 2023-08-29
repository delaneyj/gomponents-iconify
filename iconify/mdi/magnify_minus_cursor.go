package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnifyMinusCursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4a7 7 0 0 1 7 7c0 1.5-.5 3-1.39 4.19l.81.81H18l5 5l-2 2l-5-5v-.59l-.81-.81c-3.09 2.32-7.48 1.69-9.8-1.4c-2.32-3.09-1.69-7.48 1.4-9.8C8 4.5 9.5 4 11 4m-4 6v2h8v-2H7M1 1v7l7-7H1Z"/>`),
		g.Group(children),
	)
}
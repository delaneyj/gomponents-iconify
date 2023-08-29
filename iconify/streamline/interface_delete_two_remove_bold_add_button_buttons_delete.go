package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceDeleteTwoRemoveBoldAddButtonButtonsDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.19 3.05a1.06 1.06 0 0 0 0-1.49l-.75-.75a1.06 1.06 0 0 0-1.44 0L7 4.76L3.05.81a1.06 1.06 0 0 0-1.49 0l-.75.75a1.06 1.06 0 0 0 0 1.49l4 4l-4 3.95a1.06 1.06 0 0 0 0 1.49l.75.75a1.06 1.06 0 0 0 1.49 0l3.95-4l4 3.95a1.06 1.06 0 0 0 1.49 0l.75-.75a1.06 1.06 0 0 0 0-1.49L9.24 7Z"/>`),
		g.Group(children),
	)
}
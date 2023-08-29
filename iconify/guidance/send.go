package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Send(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M5.5 13L18 6m-1.75 17.5h.25a72.694 72.694 0 0 1 6.504-21.962L23.26 1L23 .74l-.538.256A72.692 72.692 0 0 1 .5 7.5v.25l5 5v7.75h.25l1.774-1.69a11.997 11.997 0 0 1 2.313-1.723L16.25 23.5Z"/>`),
		g.Group(children),
	)
}
package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountPinBoxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 21l-2 2l-2-2H4.995A1.995 1.995 0 0 1 3 19.005V4.995C3 3.893 3.893 3 4.995 3h14.01C20.107 3 21 3.893 21 4.995v14.01A1.995 1.995 0 0 1 19.005 21H14Zm-7.643-3h11.49a6.992 6.992 0 0 0-5.745-3a6.992 6.992 0 0 0-5.745 3ZM12 13a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Z"/>`),
		g.Group(children),
	)
}
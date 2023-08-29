package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5.5V0a.5.5 0 0 0-.5.5h.5Zm0 13H1a.5.5 0 0 0 .5.5v-.5ZM4 0v15h1V0H4ZM1.5 1h10V0h-10v1ZM13 2.5v9h1v-9h-1ZM11.5 13h-10v1h10v-1Zm-9.5.5V.5H1v13h1Zm11-2a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5h-1ZM11.5 1A1.5 1.5 0 0 1 13 2.5h1A2.5 2.5 0 0 0 11.5 0v1ZM7 5h4V4H7v1Z"/>`),
		g.Group(children),
	)
}
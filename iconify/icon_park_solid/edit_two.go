package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEditTwo0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M42 26v14a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h14"/><path fill="#fff" d="M14 26.72V34h7.317L42 13.308L34.695 6L14 26.72Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEditTwo0)"/>`),
		g.Group(children),
	)
}
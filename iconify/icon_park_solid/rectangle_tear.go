package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleTear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRectangleTear0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M6 8v32a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2Z"/><path stroke="#000" d="m27 6l-6 6l6 6l-6 6l6 6l-6 6l6 6"/><path stroke="#fff" d="M18 6h16M18 42h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRectangleTear0)"/>`),
		g.Group(children),
	)
}
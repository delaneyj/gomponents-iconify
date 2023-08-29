package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleTear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTRectangleTear0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M6 8v32a2 2 0 0 0 2 2h32a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2Z"/><path d="m27 6l-6 6l6 6l-6 6l6 6l-6 6l6 6M18 6h16M18 42h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTRectangleTear0)"/>`),
		g.Group(children),
	)
}
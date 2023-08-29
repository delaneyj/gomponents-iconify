package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RectangleSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSRectangleSmall0"><path fill="#fff" stroke="#fff" stroke-width="4" d="M36 14H12a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSRectangleSmall0)"/>`),
		g.Group(children),
	)
}
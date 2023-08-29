package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTShareTwo0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m26 4l18 18l-18 17V28C12 28 6 43 6 43c0-17 5-28 20-28V4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTShareTwo0)"/>`),
		g.Group(children),
	)
}
package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesAndTriangles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTCirclesAndTriangles0"><g fill="none" stroke="#fff" stroke-width="4"><path d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="m24 4l17.32 30H6.68L24 4Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTCirclesAndTriangles0)"/>`),
		g.Group(children),
	)
}
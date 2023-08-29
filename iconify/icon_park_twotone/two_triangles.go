package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoTriangles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTwoTriangles0"><g fill="none"><path fill="#555" d="m24 4l17.32 30H6.68L24 4Z"/><path fill="#555" d="m24 44l17.32-30H6.68L24 44Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 4l17.32 30H6.68L24 4Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 44l17.32-30H6.68L24 44Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTwoTriangles0)"/>`),
		g.Group(children),
	)
}
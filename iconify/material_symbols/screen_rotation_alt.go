package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotationAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.175 21.4L3.75 13H6.6l7 7l5-5H16v-2h6v6h-2v-2.6l-5 5q-.575.575-1.413.575t-1.412-.575ZM2 11V5h2v2.6l5-5q.575-.575 1.413-.575t1.412.575L20.25 11H17.4l-7-7l-5 5H8v2H2Z"/>`),
		g.Group(children),
	)
}
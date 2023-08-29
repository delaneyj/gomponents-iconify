package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotationAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.175 21.4L3.75 13H6.6l7 7l5-5H17q-.425 0-.713-.288T16 14q0-.425.288-.713T17 13h4q.425 0 .713.288T22 14v4q0 .425-.288.713T21 19q-.425 0-.713-.288T20 18v-1.6l-5 5q-.575.575-1.413.575t-1.412-.575ZM3 11q-.425 0-.713-.288T2 10V6q0-.425.288-.713T3 5q.425 0 .713.288T4 6v1.6l5-5q.575-.575 1.413-.575t1.412.575L20.25 11H17.4l-7-7l-5 5H7q.425 0 .713.288T8 10q0 .425-.288.713T7 11H3Z"/>`),
		g.Group(children),
	)
}
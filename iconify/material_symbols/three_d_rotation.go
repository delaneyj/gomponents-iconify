package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeDRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12h2q0 2.875 1.813 5.075t4.637 2.775L9 18.4l1.4-1.4l4.55 4.55q-.725.25-1.463.35T12 22Zm.5-7V9h3q.425 0 .713.288T16.5 10v4q0 .425-.288.713T15.5 15h-3Zm-5 0v-1.5H10v-1H8.5v-1H10v-1H7.5V9h3q.425 0 .713.288T11.5 10v4q0 .425-.288.713T10.5 15h-3Zm6.5-1.5h1v-3h-1v3Zm6-1.5q0-2.875-1.813-5.075T13.55 4.15L15 5.6L13.6 7L9.05 2.45q.725-.25 1.463-.35T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12h-2Z"/>`),
		g.Group(children),
	)
}
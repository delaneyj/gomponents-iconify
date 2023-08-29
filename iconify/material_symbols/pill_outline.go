package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.625 21q-2.35 0-3.988-1.638T3 15.376q0-1.125.425-2.15T4.65 11.4l6.75-6.75q.8-.8 1.825-1.225T15.375 3q2.35 0 3.988 1.637T21 8.625q0 1.125-.425 2.15T19.35 12.6l-6.75 6.75q-.8.8-1.825 1.225T8.625 21Zm6.65-7.15l2.675-2.65q.5-.5.775-1.175t.275-1.4q0-1.5-1.063-2.563T15.375 5q-.725 0-1.4.275T12.8 6.05l-2.65 2.675l5.125 5.125ZM8.625 19q.725 0 1.4-.275t1.175-.775l2.65-2.675l-5.125-5.125L6.05 12.8q-.5.5-.775 1.175T5 15.375q0 1.5 1.063 2.563T8.624 19Z"/>`),
		g.Group(children),
	)
}
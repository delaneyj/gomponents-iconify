package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.35 14.6l-1.4-1.4l2-2q.5-.5.775-1.175t.275-1.4q0-1.5-1.063-2.563T15.376 5q-.725 0-1.4.275T12.8 6.05l-2 2l-1.4-1.4l2-2q.8-.8 1.825-1.225T15.375 3q2.35 0 3.987 1.638T21 8.625q0 1.125-.425 2.15T19.35 12.6l-2 2ZM8.625 19q.725 0 1.4-.275t1.175-.775l1.925-1.925l-5.15-5.15L6.05 12.8q-.5.5-.775 1.175T5 15.375q0 1.5 1.063 2.563T8.625 19ZM19.8 22.7l-5.275-5.275L12.6 19.35q-.8.8-1.825 1.225T8.625 21q-2.35 0-3.988-1.637T3 15.375q0-1.125.425-2.15T4.65 11.4l1.925-1.925L1.4 4.3l1.4-1.4l18.4 18.4l-1.4 1.4Zm-9.25-9.25Zm2.825-2.825Z"/>`),
		g.Group(children),
	)
}
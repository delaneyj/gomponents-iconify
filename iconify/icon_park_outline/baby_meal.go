package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyMeal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 6H10a5 5 0 0 0-5 5v26a5 5 0 0 0 5 5h20"/><path stroke-linecap="round" stroke-linejoin="round" d="M19 6v17.524L5 28m14-4l12 4"/><path stroke-linecap="round" d="M37 18v24"/><path d="M31 13.333C31 7.111 35 4 37 4s6 3.111 6 9.333c0 6.223-12 6.223-12 0Z"/></g>`),
		g.Group(children),
	)
}
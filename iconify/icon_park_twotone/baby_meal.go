package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyMeal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBabyMeal0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M24 6H10a5 5 0 0 0-5 5v26a5 5 0 0 0 5 5h20"/><path stroke-linecap="round" stroke-linejoin="round" d="M19 6v17.524L5 28m14-4l12 4"/><path stroke-linecap="round" d="M37 18v24"/><path fill="#555" d="M31 13.333C31 7.111 35 4 37 4s6 3.111 6 9.333c0 6.223-12 6.223-12 0Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBabyMeal0)"/>`),
		g.Group(children),
	)
}
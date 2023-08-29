package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeHundredSixtyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.3 19.3q-.275-.275-.275-.7t.275-.7l1.1-1.1q-3.2-.425-5.3-1.75T2 12q0-2.075 2.888-3.538T12 7q4.225 0 7.113 1.463T22 12q0 1.35-1.3 2.475t-3.475 1.8q-.5.15-.863-.125T16 15.325q0-.3.213-.587t.512-.388q1.575-.5 2.425-1.175T20 12q0-.8-2.137-1.9T12 9q-3.725 0-5.863 1.1T4 12q0 .6 1.275 1.438T8.9 14.7l-.6-.6q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l2.6 2.6q.15.15.212.325t.063.375q0 .2-.063.375t-.212.325l-2.6 2.6q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}
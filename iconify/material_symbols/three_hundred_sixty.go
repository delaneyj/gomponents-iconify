package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreeHundredSixty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9 20l-1.4-1.4l1.75-1.8q-3.2-.425-5.275-1.75T2 12q0-2.075 2.888-3.538T12 7q4.225 0 7.113 1.463T22 12q0 1.55-1.663 2.775T16 16.6v-2.05q1.925-.5 2.963-1.238T20 12q0-.8-2.138-1.9T12 9q-3.725 0-5.863 1.1T4 12q0 .6 1.275 1.438T8.9 14.7l-1.3-1.3L9 12l4 4l-4 4Z"/>`),
		g.Group(children),
	)
}
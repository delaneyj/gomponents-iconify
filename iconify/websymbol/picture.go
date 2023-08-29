package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1360 1v1000H0V1h1360zM120 881h1120V121H120v760zm1000-560q0 33-23.5 56.5T1040 401t-56.5-23.5T960 321t23.5-56.5T1040 241t56.5 23.5T1120 321zm43 473q2 1 1 4t-4 3H204q-4 0-4-4v-2q18-36 135-280t144-299q1-2 3.5-2.5t3.5 1.5l311 386l152-122h7l1 2q39 50 101.5 155T1163 794z"/>`),
		g.Group(children),
	)
}
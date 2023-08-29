package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dialog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M695 5q20 0 33 13t13 33v394q0 19-13 32t-33 14H463v162L278 491H46q-19 0-32-14T0 445V51q0-20 14-33T46 5h649z"/>`),
		g.Group(children),
	)
}
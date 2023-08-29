package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M718 191q10 0 17 7t6 17v92H0v-92q0-10 7-17t16-7h328L530 13q7-7 17-7t16 7l33 33q7 7 7 16t-7 16L483 191h235zM46 354h649l-43 258q-2 8-8 14t-15 6H112q-8 0-14-6t-9-14z"/>`),
		g.Group(children),
	)
}
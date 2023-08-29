package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartAreaMulti(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2h2Zm17 2.086V18.21H6v-6.152l6.59-5.99l2.966 3.461L21 4.086ZM8.258 16.21H19v-3.796l-3.73 3.73l-3.034-3.539l-3.978 3.605ZM19 9.586v-.672l-3.556 3.557l-3.033-3.538L8 12.943v.801l4.415-4.001l2.966 3.46L19 9.586Z"/>`),
		g.Group(children),
	)
}
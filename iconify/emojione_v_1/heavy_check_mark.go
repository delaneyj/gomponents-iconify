package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyCheckMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#405866" d="M60.59 11.873a6.58 6.58 0 0 0-9.165 1.621l-25.02 35.732L10 38.02a6.01 6.01 0 0 0-6.787 9.922l21.99 15.02a5.96 5.96 0 0 0 4.682.889a6.579 6.579 0 0 0 4.257-2.708l28.07-40.1a6.59 6.59 0 0 0-1.623-9.173"/>`),
		g.Group(children),
	)
}
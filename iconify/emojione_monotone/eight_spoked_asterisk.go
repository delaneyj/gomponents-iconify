package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightSpokedAsterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m62 32l-24.166-2.417l15.379-18.797l-18.797 15.38L32 2l-2.417 24.166l-18.797-15.38l15.38 18.797L2 32l24.166 2.416l-15.38 18.797l18.797-15.379L32 62l2.416-24.166l18.797 15.379l-15.379-18.797z"/>`),
		g.Group(children),
	)
}
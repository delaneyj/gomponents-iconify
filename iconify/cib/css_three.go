package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m30 2l-4 23l-13.956 5L0 25l1.225-5.925H6.35l-.5 2.538l7.275 2.775l8.381-2.775l1.175-6.069H1.844l1-5.125H23.7l.656-3.294H3.519L4.538 2z"/>`),
		g.Group(children),
	)
}
package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Addthis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 13.496h-4.501v4.484h-3v-4.484H6v-2.99h4.5V6.021h3.001v4.485H18v2.99zM21 .041H3a3.004 3.004 0 0 0-3 2.99v17.94a3.004 3.004 0 0 0 3 2.988h18a3.005 3.005 0 0 0 3-2.988V3.031a3.005 3.005 0 0 0-3-2.99z"/>`),
		g.Group(children),
	)
}
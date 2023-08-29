package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LandRowsHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 20V4c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v16c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2M4 6.5V4h16v2.5H4M4 11V8.5h16V11H4m0 4.5V13h16v2.5H4M4 20v-2.5h16V20H4Z"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.293 3L11.146.854l.708-.708l3 3a.5.5 0 0 1 0 .708l-3 3l-.708-.708L13.293 4H3.5A2.5 2.5 0 0 0 1 6.5V8H0V6.5A3.5 3.5 0 0 1 3.5 3h9.793ZM15 7v1.5a3.5 3.5 0 0 1-3.5 3.5H1.707l2.147 2.146l-.708.708l-3-3a.5.5 0 0 1 0-.707l3-3l.708.707L1.707 11H11.5A2.5 2.5 0 0 0 14 8.5V7h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
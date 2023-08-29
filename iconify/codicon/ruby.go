package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m1 7.19l6.64 6.64h.72L15 7.19v-.72l-3.32-3.32l-.36-.15H4.68l-.36.15L1 6.47v.72zm7 5.56L2.08 6.83L4.89 4h6.22l2.81 2.83L8 12.75zm0-7.73h2.69l1.81 1.81l-4.5 4.4V5.02z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
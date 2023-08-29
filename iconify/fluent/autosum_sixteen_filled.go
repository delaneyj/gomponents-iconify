package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutosumSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.06 1.956a.75.75 0 0 1 .69-.456h9.5a.75.75 0 0 1 0 1.5H4.51l3.915 4.08a.75.75 0 0 1 .014 1.023L4.445 12.5h7.805a.75.75 0 0 1 0 1.5h-9.5a.75.75 0 0 1-.555-1.254l4.663-5.133l-4.65-4.844a.75.75 0 0 1-.148-.813Z"/>`),
		g.Group(children),
	)
}
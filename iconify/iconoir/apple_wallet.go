package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppleWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2Z"/><path d="M3 15h6.4c.331 0 .605.278.75.576c.206.423.694.924 1.85.924c1.156 0 1.644-.5 1.85-.924c.145-.298.419-.576.75-.576H21M3 7h18M3 11h18"/></g>`),
		g.Group(children),
	)
}
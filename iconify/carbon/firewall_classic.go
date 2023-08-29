package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirewallClassic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 20v-3a4 4 0 0 0-8 0v3a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h8a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2zm-6-3a2 2 0 0 1 4 0v3h-4zm6 11h-8v-6h8zm-13-1H4a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h11v2H4v3h11zm2-9H8a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h11v2H8v3h9zm5-9H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2zM4 7h18V4H4z"/>`),
		g.Group(children),
	)
}
package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VlanIbm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 22a4 4 0 1 0 4 4a4.005 4.005 0 0 0-4-4zm0 6a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zM30 5a3 3 0 1 0-4 2.815V15h-9V9h-2v6H6V7.816a3 3 0 1 0-2 0V15a2.002 2.002 0 0 0 2 2h9v3h2v-3h9a2.002 2.002 0 0 0 2-2V7.816A2.996 2.996 0 0 0 30 5zM5 4a1 1 0 1 1-1 1a1.001 1.001 0 0 1 1-1zm22 2a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1z"/><circle cx="16" cy="5" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}
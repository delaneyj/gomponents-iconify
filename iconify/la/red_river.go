package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedRiver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm5 4a1 1 0 0 0-1 1v6a2 2 0 0 0 2-2v-2a1 1 0 0 1 1-1h2a2 2 0 0 0 2-2h-6zm4 4a1 1 0 0 0-1 1v6a2 2 0 0 0 2-2v-2a1 1 0 0 1 1-1h2a2 2 0 0 0 2-2h-6z"/>`),
		g.Group(children),
	)
}
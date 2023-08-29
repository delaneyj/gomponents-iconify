package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v6h19V5H5zm2 2h15v2H7V7zm2 6v6h19v-6H9zm2 2h15v2H11v-2zm-6 6v6h19v-6H5zm2 2h15v2H7v-2z"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.207 5.5L5.707 7H11v1H5.707l1.5 1.5l-.707.707L3.793 7.5L6.5 4.793l.707.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
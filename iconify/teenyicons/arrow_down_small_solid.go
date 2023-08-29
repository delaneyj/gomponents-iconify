package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 4v5.293l1.5-1.5l.707.707L7.5 11.207L4.793 8.5l.707-.707l1.5 1.5V4h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
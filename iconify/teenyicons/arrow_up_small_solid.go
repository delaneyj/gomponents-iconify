package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 3.793L10.207 6.5l-.707.707l-1.5-1.5V11H7V5.707l-1.5 1.5l-.707-.707L7.5 3.793Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
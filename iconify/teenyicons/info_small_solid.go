package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoSmallSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3.99V5H7V3.99h1ZM6 11v-1h1V8H6V7h2v3h1v1H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
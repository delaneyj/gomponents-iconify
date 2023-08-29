package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainProfile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 25H2v2h2v2h2v-2h5v2h2v-2h5v2h2v-2h5v2h2v-2h3v-2zM8 16H2v-2h6v-2H2v-2h6a2.002 2.002 0 0 1 2 2v2a2.002 2.002 0 0 1-2 2z"/><path fill="currentColor" d="m28.55 14.23l-8.58-7.864A8.977 8.977 0 0 0 13.888 4H2v2h10v4a2.002 2.002 0 0 0 2 2h9.156l4.042 3.705A2.472 2.472 0 0 1 25.528 20H2v2h23.527a4.473 4.473 0 0 0 3.023-7.77ZM14 10V6.005a6.977 6.977 0 0 1 4.618 1.835L20.975 10Z"/>`),
		g.Group(children),
	)
}
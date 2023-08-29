package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimePlot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.586 13L21 10.414V6h2v3.586l2 2L23.586 13z"/><path fill="currentColor" d="M22 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8Zm0-14a6 6 0 1 0 6 6a6.007 6.007 0 0 0-6-6Z"/><path fill="currentColor" d="m8.63 18l7 6H30v-2H16.37l-7-6H4V2H2v26a2.002 2.002 0 0 0 2 2h26v-2H4V18Z"/>`),
		g.Group(children),
	)
}
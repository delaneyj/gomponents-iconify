package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 128h405v43H0v-43zm0 128v-43h107v43H0zm149 0v-43h107v43H149zm150 0v-43h106v43H299z"/>`),
		g.Group(children),
	)
}
package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeUnfolded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 15h10v5h5v-5h5v-5H12V5H7v5H2z"/><path d="M7 15v-5h5v5h5v-5"/></g>`),
		g.Group(children),
	)
}
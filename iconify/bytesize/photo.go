package bytesize

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20 24l-8-8L2 26V2h28v22m-14-4l6-6l8 8v8H2v-6"/><circle cx="10" cy="9" r="3"/></g>`),
		g.Group(children),
	)
}
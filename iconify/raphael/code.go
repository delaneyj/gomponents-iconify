package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8.982 7.107L.322 15.77l8.66 8.662l3.15-3.15l-5.51-5.512l5.51-5.51l-3.15-3.153zm12.675 0l-3.148 3.15l5.51 5.512l-5.51 5.51l3.147 3.15l8.662-8.662l-8.663-8.66z"/>`),
		g.Group(children),
	)
}
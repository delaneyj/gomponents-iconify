package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoofingContractor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M8.149 16.062h6.06l-.053 3.575l-6.007 5.406v-8.981zm16.758-1.979L1 35.169L3.52 38l21.485-18.954L46.486 38L49 35.169L25.097 14.083L25 14l-.093.083z"/>`),
		g.Group(children),
	)
}
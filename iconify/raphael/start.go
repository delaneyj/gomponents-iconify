package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Start(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.316 5.318L9.833 13.682V5.5H5.5v20h4.333v-8.182l14.483 8.364z"/>`),
		g.Group(children),
	)
}
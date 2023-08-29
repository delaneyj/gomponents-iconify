package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 24.001v-3.165h27.693v3.165zm0-7.912v-3.165h27.693v3.165zm0-7.055V.461A.505.505 0 0 1 .748.059L.745.058l8.483 4.273a.464.464 0 0 1 .003.836l-.003.001L.745 9.437a.534.534 0 0 1-.242.059h-.02A.484.484 0 0 1 0 9.035v-.001zm12.659-2.44V3.429h15.033v3.165z"/>`),
		g.Group(children),
	)
}
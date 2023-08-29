package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h4v1H2.707l3.147 3.146l-.708.708L2 2.707V5H1V1Zm11.293 1H10V1h4v4h-1V2.707L9.854 5.854l-.708-.708L12.293 2Zm-6.44 7.854L2.708 13H5v1H1v-4h1v2.293l3.146-3.147l.708.708Zm4-.708L13 12.293V10h1v4h-4v-1h2.293L9.146 9.854l.708-.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
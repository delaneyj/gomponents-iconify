package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.293 2H10V1h4v4h-1V2.707L9.854 5.854l-.708-.708L12.293 2Zm-6.44 7.854L2.708 13H5v1H1v-4h1v2.293l3.146-3.147l.708.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimiseAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L10.707 5H13v1H9V2h1v2.293l3.146-3.147l.708.708ZM2 9h4v4H5v-2.293l-3.146 3.147l-.708-.707L4.293 10H2V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
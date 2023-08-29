package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinimiseSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L10.707 5H13v1H9V2h1v2.293l3.146-3.147l.708.708ZM4.293 5L1.146 1.854l.708-.708L5 4.293V2h1v4H2V5h2.293ZM2 9h4v4H5v-2.293l-3.146 3.147l-.708-.707L4.293 10H2V9Zm7 0h4v1h-2.293l3.147 3.146l-.708.708L10 10.707V13H9V9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DollarSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 1V0h1v1h1.5A3.5 3.5 0 0 1 13 4.5h-1A2.5 2.5 0 0 0 9.5 2h-4a2.5 2.5 0 0 0 0 5h4a3.5 3.5 0 1 1 0 7H8v1H7v-1H5.5A3.5 3.5 0 0 1 2 10.5h1A2.5 2.5 0 0 0 5.5 13h4a2.5 2.5 0 0 0 0-5h-4a3.5 3.5 0 1 1 0-7H7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
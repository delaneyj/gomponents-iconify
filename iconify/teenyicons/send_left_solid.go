package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendLeftSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 1h1v13h-1V1ZM2.707 8l3.147 3.146l-.708.708L.793 7.5l4.353-4.354l.708.708L2.707 7H12v1H2.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
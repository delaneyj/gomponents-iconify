package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.5 0a6.5 6.5 0 1 0 4.23 11.436l3.416 3.418l.708-.708l-3.418-3.417A6.5 6.5 0 0 0 6.5 0Z"/>`),
		g.Group(children),
	)
}
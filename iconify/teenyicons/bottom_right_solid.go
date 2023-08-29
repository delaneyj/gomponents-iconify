package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomRightSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.854 1.146L13 12.293V8h1v6H8v-1h4.293L1.146 1.854l.708-.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
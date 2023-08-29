package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomLeftSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.854 1.854L2.707 13H7v1H1V8h1v4.293L13.146 1.146l.708.708Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}
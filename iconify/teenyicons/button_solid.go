package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4 6a3 3 0 0 1 0-6h7a3 3 0 1 1 0 6H9V3.5a2.5 2.5 0 0 0-5 0V6Z"/><path fill="currentColor" d="M6.5 2A1.5 1.5 0 0 0 5 3.5v4.55a2.5 2.5 0 0 0-2 2.45A4.5 4.5 0 0 0 7.5 15H8a5 5 0 0 0 5-5v-.853A2.147 2.147 0 0 0 10.853 7H8V3.5A1.5 1.5 0 0 0 6.5 2Z"/>`),
		g.Group(children),
	)
}
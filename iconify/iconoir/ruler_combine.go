package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCombine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 21.4V2.6a.6.6 0 0 1 .6-.6h18.8a.6.6 0 0 1 .6.6v6.8a.6.6 0 0 1-.6.6H10.6a.6.6 0 0 0-.6.6v10.8a.6.6 0 0 1-.6.6H2.6a.6.6 0 0 1-.6-.6ZM16 10V7m-6 3V7m0 9H7m3-6H7"/>`),
		g.Group(children),
	)
}
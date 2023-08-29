package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrewDriver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m14.146 12.334l-1.83-.599l-4.339-4.338l-.579.58l4.338 4.337l.637 1.872l3.067 1.832l.578-.579l-1.872-3.105zM8.255 4.999L7.209 6.054c.157-.39.106-.819-.171-1.101L2.423.281C2.035-.11 1.356-.061.904.396L.393.912c-.45.457-.5 1.145-.111 1.537L4.896 7.12c.276.28.701.33 1.085.173l-.947.955l.715.72l3.22-3.249l-.714-.72z"/>`),
		g.Group(children),
	)
}
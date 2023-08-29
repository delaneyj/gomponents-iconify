package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m.511 13.019l16.695 10.836c.656.427 1.461-.136 1.461-1.02V16.15l11.872 7.705c.656.427 1.461-.136 1.461-1.02V1.166c0-.884-.805-1.447-1.461-1.02L18.667 7.85V1.166c0-.884-.805-1.447-1.461-1.02L.511 10.981a1.27 1.27 0 0 0-.003 2.037l.003.002z"/>`),
		g.Group(children),
	)
}
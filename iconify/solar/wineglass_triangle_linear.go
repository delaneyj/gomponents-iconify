package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineglassTriangleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m12 14.571l8.516-8.707C21.562 4.794 20.8 3 19.3 3H4.701C3.2 3 2.438 4.794 3.484 5.864L12 14.57Zm0 0V21m0 0h4.244M12 21H7.756M7.473 9.75h9.054"/>`),
		g.Group(children),
	)
}
package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 5a4.514 4.514 0 0 1 4.481 4.154l.141 1.861l1.867-.013l.093-.001A3.003 3.003 0 0 1 21 14c0 .748-.28 1.452-.783 2H3.279A1.906 1.906 0 0 1 3 15c0-1.074.851-1.953 1.915-1.998c.059.007.118.012.178.015l2.659.124l-.621-2.588A4.468 4.468 0 0 1 7 9.5C7 7.019 9.019 5 11.5 5m0-2A6.5 6.5 0 0 0 5 9.5a6.5 6.5 0 0 0 .186 1.519C5.123 11.016 5.064 11 5 11a4 4 0 0 0-4 4c0 1.202.541 2.267 1.38 3h18.593C22.196 17.089 23 15.643 23 14a5 5 0 0 0-5-5l-.025.002A6.496 6.496 0 0 0 11.5 3z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMultipleTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 4a5.5 5.5 0 0 0-1.99 10.63c.04-.7.19-1.37.433-1.994A3.5 3.5 0 0 1 6.5 6h6a3.5 3.5 0 1 1 0 7H12a1 1 0 1 0 0 2h.5a5.5 5.5 0 1 0 0-11h-6Zm12.613 8.213A3.001 3.001 0 0 1 18 18h-7a3 3 0 1 1 0-6h1a1 1 0 0 0 0-2h-1a5 5 0 0 0 0 10h7a5 5 0 0 0 1.496-9.772a6.48 6.48 0 0 1-.383 1.985Z"/>`),
		g.Group(children),
	)
}
package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.659 13.593a2.25 2.25 0 0 1 0-3.182l7.752-7.752a2.25 2.25 0 0 1 3.182 0l7.751 7.752a2.25 2.25 0 0 1 0 3.182l-7.751 7.751a2.25 2.25 0 0 1-3.182 0l-7.752-7.751Z"/>`),
		g.Group(children),
	)
}
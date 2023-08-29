package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRadiusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10 1.5h1m2 0h1m-1 3h1m-1 3h1m-1 3h1m-1 3h1m-4 0h1m-4 0h1m-4 0h1m-4 0h1m-1-3h1M1.5 8V5.5a4 4 0 0 1 4-4H8"/>`),
		g.Group(children),
	)
}
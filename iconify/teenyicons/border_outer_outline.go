package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderOuterOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7 4.5h1m2 3h1m-4 0h1m-1 3h1m-4-3h1m-3.5-6h12v12h-12v-12Z"/>`),
		g.Group(children),
	)
}
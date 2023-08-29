package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DollarOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M2.5 10.5a3 3 0 0 0 3 3h4a3 3 0 1 0 0-6h-4a3 3 0 0 1 0-6h4a3 3 0 0 1 3 3M7.5 0v1.5m0 13.5v-1.5"/>`),
		g.Group(children),
	)
}
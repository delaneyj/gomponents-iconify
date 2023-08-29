package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImacSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M13.5 0A1.5 1.5 0 0 1 15 1.5V7H0V1.5A1.5 1.5 0 0 1 1.5 0h12ZM0 8v1.5A1.5 1.5 0 0 0 1.5 11H7v3H2v1h11v-1H8v-3h5.5A1.5 1.5 0 0 0 15 9.5V8H0Z"/>`),
		g.Group(children),
	)
}
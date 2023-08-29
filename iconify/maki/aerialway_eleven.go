package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AerialwayEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9 4.5H6V3.1a1 1 0 0 0 .43-.52L9.5 2a.5.5 0 0 0 0-1l-3.25.61a1 1 0 0 0-1.69.32L1.5 2.5a.5.5 0 0 0 0 1l3.25-.61A1 1 0 0 0 5 3.1v1.4H2a1 1 0 0 0-1 1V9a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V5.5a1 1 0 0 0-1-1zm-4 4H2.5v-3H5v3zm3.5 0H6v-3h2.5v3z" fill="currentColor"/>`),
		g.Group(children),
	)
}
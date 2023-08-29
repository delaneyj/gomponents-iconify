package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElbowConnectorSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1.5 0a1.5 1.5 0 1 0 1.415 2H7v12h5.085a1.5 1.5 0 1 0 0-1H8V1H2.915A1.5 1.5 0 0 0 1.5 0Z"/>`),
		g.Group(children),
	)
}
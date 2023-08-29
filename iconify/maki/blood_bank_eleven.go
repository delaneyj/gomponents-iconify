package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BloodBankEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8.405 4.644h.001L5.5 1L2.593 4.644h.002A3.37 3.37 0 0 0 2 6.56a3.464 3.464 0 0 0 3.5 3.43A3.464 3.464 0 0 0 9 6.558a3.37 3.37 0 0 0-.595-1.915zM8 7H6v2H5V7H3V6h2V4h1v2h2v1z" fill="currentColor"/>`),
		g.Group(children),
	)
}
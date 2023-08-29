package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsuranceAgency(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25.511 2.892L1.268 23H7v24h36V23h6.268L25.511 2.892zM30.942 40H19.593l2.38-10.711c-1.439-1-2.38-2.506-2.38-4.35c0-3.038 2.541-5.439 5.674-5.439c3.135 0 5.675 2.493 5.675 5.531c0 1.845-.941 3.245-2.379 4.242L30.942 40z"/>`),
		g.Group(children),
	)
}
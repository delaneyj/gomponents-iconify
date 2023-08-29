package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Email(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M695 1q20 0 33 13t13 33v464q0 19-13 32t-33 14H46q-19 0-32-14T0 511V47q0-19 14-33T46 1h649zm-53 97q2-1 2-3t-3-1H101q-2 0-3 1t1 3l259 172q13 10 25 0z"/>`),
		g.Group(children),
	)
}
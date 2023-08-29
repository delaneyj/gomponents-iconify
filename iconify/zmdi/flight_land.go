package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightLand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 368h405v43H0v-43zm153-122L40 215l-34-9V96l31 8l20 50l106 28V5l41 11l59 193l113 30q13 3 19.5 14.5t3 24.5t-15 19.5T359 301l-113-30z"/>`),
		g.Group(children),
	)
}
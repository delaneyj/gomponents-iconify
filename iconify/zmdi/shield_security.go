package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldSecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m192 5l192 86v128q0 89-55 162.5T192 475q-82-20-137-93.5T0 219V91zm0 235V52L43 118v122h149v191q59-19 100-72t49-119H192z"/>`),
		g.Group(children),
	)
}
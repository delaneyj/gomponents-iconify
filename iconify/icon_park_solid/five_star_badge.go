package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiveStarBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFiveStarBadge0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M23.103 20.817a1 1 0 0 1 1.794 0l2.985 6.048a1 1 0 0 0 .753.548l6.675.97a1 1 0 0 1 .554 1.705l-4.83 4.708a1 1 0 0 0-.288.885l1.14 6.648a1 1 0 0 1-1.45 1.054l-5.97-3.138a1 1 0 0 0-.931 0l-5.97 3.138a1 1 0 0 1-1.452-1.054l1.14-6.648a1 1 0 0 0-.287-.885l-4.83-4.708a1 1 0 0 1 .554-1.706l6.675-.97a1 1 0 0 0 .753-.547l2.985-6.048ZM36 4H12v10l12 5l12-5V4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFiveStarBadge0)"/>`),
		g.Group(children),
	)
}
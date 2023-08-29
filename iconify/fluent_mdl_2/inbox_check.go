package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InboxCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M832 794L467 429l90-90l275 275l531-531l90 90l-621 621zm1088 482q1 2 1 29t1 69t0 97t0 111t-1 114t0 102t-1 79t0 43H0v-121q0-47-1-103t0-113t-1-112t0-96t0-70t2-29l238-714l101 102l-205 616h418l128 256h560l128-256h418l-269-809l101-101l302 906zm-128 132h-344l-128 256H600l-128-256H128v384h1664v-384z"/>`),
		g.Group(children),
	)
}
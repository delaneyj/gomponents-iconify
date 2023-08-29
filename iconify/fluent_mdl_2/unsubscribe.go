package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unsubscribe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1276q1 2 1 28t1 69t0 97t0 112t-1 114t0 102t-1 79t0 43H0v-644L373 158l102 99l-341 1023h418l128 256h560l128-256h418L1445 256l101-103l374 1123zm-128 132h-344l-128 256H600l-128-256H128v384h1664v-384zM531 787l338-339l-338-339l90-90l339 339l339-339l90 90l-339 339l339 339l-90 90l-339-339l-339 339l-90-90z"/>`),
		g.Group(children),
	)
}
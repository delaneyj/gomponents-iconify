package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreadmillOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTreadmillOne0"><g fill="none" stroke="#fff" stroke-miterlimit="2" stroke-width="4"><path fill="#555" d="M30 14a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m11 21l6-5l8 3l-3 6l7 6l2 7"/><path stroke-linecap="round" stroke-linejoin="round" d="m22 25l-5 8l-8-1M6 44h34l4-9M25 19l8 5l6-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTreadmillOne0)"/>`),
		g.Group(children),
	)
}
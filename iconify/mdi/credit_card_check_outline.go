package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19c0-.34.04-.67.09-1H3v-6h16v1c.7 0 1.37.13 2 .35V6c0-1.11-.89-2-2-2H3c-1.11 0-2 .89-2 2v12a2 2 0 0 0 2 2h10.09c-.05-.33-.09-.66-.09-1M3 6h16v2H3V6m14.75 16L15 19l1.16-1.16l1.59 1.59l3.59-3.59l1.16 1.41L17.75 22"/>`),
		g.Group(children),
	)
}
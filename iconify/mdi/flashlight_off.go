package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 5.27L3.28 4L20 20.72L18.73 22L15 18.27V22H9v-9.73l-7-7M18 5l-3 5h-3.18l-5-5H18m0-1H6V2h12v2m-3 7v2.18L12.82 11H15Z"/>`),
		g.Group(children),
	)
}
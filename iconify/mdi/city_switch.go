package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CitySwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 15V9h-6V3l-3-3l-3 3v2H3v10h18m-2-2h-2v-2h2v2m-8-8V3h2v2h-2m0 4V7h2v2h-2m0 4v-2h2v2h-2M7 9H5V7h2v2m0 4H5v-2h2v2m2 4v2h6v-2l3 3l-3 3v-2H9v2l-3-3l3-3Z"/>`),
		g.Group(children),
	)
}
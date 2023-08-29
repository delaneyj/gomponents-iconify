package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M235 5q26 0 45 19t19 45v342q0 26-19 45t-45 19H64q-27 0-45.5-19T0 411V69q0-26 18.5-45T64 5h171zm-43 427v-21h-85v21h85zm69-64V69H37v299h224z"/>`),
		g.Group(children),
	)
}
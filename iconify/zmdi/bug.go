package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M341 107v42h-44q2 13 2 22v21h42v43h-42v21q0 9-2 21h44v43h-60q-17 29-46 46.5T171 384t-64.5-17.5T60 320H0v-43h45q-2-12-2-21v-21H0v-43h43v-21q0-9 2-22H0v-42h60q15-26 39-42L64 30L94 0l47 46q14-3 29.5-3t30.5 3l46-46l30 30l-34 35q24 16 38 42h60zM213 277v-42h-85v42h85zm0-85v-43h-85v43h85z"/>`),
		g.Group(children),
	)
}
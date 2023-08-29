package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookSettingsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 0H6C4.89 0 4 .895 4 2v16c0 1.11.89 2 2 2h12c1.11 0 2-.89 2-2V2c0-1.105-.89-2-2-2m0 18H6V2h2v8l2.5-2.25L13 10V2h5v16M7 22h2v2H7v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}
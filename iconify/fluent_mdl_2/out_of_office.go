package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutOfOffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1600 1024q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45q0-26 19-45t45-19zm-832 256h128v768H768v-768zM1920 0v2048h-128V128H896v768H768V0h1152zM347 1024h805v128H347l210 211l-90 90l-365-365l365-365l90 90l-210 211z"/>`),
		g.Group(children),
	)
}
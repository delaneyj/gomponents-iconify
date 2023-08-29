package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 384v448q0 83-29 158t-80 135t-122 98t-153 52v773H896v-773q-83-12-153-51t-121-99t-81-134t-29-159V384h128V0h128v384h384V0h128v384h128zm-128 128H640v320q0 66 25 124t68 102t102 69t125 25q66 0 124-25t101-68t69-102t26-125V512z"/>`),
		g.Group(children),
	)
}
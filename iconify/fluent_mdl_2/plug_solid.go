package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 384v448q0 84-29 159t-80 134t-122 98t-153 52v773H896v-773q-83-12-153-51t-121-99t-81-134t-29-159V384h128V0h128v384h384V0h128v384h128z"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConstructionCone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1650 1920h270v128H128v-128h270L846 128h356l448 1792zM626 1536l-96 384h988L1102 256H946l-192 768h414l32 128H722l-64 256h606l32 128H626z"/>`),
		g.Group(children),
	)
}
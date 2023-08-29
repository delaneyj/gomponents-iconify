package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1280 384H640V256h640v128zm0 1024H640v-128h640v128zm0 256H640v-128h640v128zM1408 0q27 0 50 10t40 27t28 41t10 50v1792H384V128q0-27 10-50t27-40t41-28t50-10h896zm0 128H512v1664h896V128z"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1792H0V128zm1920 128H128v256h1792V256zM128 1792h1792V640H128v1152zm128-640V768h384v384H256zm128-256v128h128V896H384zm-128 768v-384h384v384H256zm128-256v128h128v-128H384zm512-384V896h768v128H896zm0 512v-128h768v128H896z"/>`),
		g.Group(children),
	)
}
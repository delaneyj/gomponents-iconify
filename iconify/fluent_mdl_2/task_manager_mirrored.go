package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskManagerMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1792H0V128zm1920 128H128v256h1792V256zM128 1792h1792V640H128v1152zm1280-640V768h384v384h-384zm128-256v128h128V896h-128zm-128 768v-384h384v384h-384zm128-256v128h128v-128h-128zM384 1024V896h768v128H384zm0 512v-128h768v128H384z"/>`),
		g.Group(children),
	)
}
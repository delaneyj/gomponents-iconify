package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FocusView(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 640V256h128v512H256V640h384zm-384 768v-128h512v512H640v-384H256zm1024 384v-512h512v128h-384v384h-128zm128-1152h384v128h-512V256h128v384z"/>`),
		g.Group(children),
	)
}
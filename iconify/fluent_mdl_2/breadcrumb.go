package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Breadcrumb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 512h1536v128H256V512zm384 512V896h1152v128H640zm384 384v-128h768v128h-768z"/>`),
		g.Group(children),
	)
}
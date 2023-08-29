package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1536v-128H768v128h1280zm0-640H0v128h2048V896zM384 384v128h1664V384H384z"/>`),
		g.Group(children),
	)
}
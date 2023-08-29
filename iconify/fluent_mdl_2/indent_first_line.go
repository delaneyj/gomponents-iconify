package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentFirstLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 128h768v128h-768V128zM768 1792v-128h1152v128H768zm0-1024V640h1152v128H768zm0 512v-128h1152v128H768zM573 3l317 317l-317 317l-90-90l163-163H0V256h646L483 93l90-90z"/>`),
		g.Group(children),
	)
}
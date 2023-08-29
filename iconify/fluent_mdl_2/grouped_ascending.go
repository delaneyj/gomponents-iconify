package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupedAscending(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 128h1536v128H384V128zm128 603v1061h1408v128H384V731L91 1024L0 933l448-447l448 447l-91 91l-293-293zm512 677v-128h896v128h-896zm0-384V896h896v128h-896zm0-384V512h896v128h-896z"/>`),
		g.Group(children),
	)
}
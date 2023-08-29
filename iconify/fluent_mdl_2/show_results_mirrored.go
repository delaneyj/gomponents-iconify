package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowResultsMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 1792v-128H0v128h1664zm384-1408V256h-128v128h128zm0 1024v-128h-128v128h128zm-384 0v-128H0v128h1664zm0-1152H0v128h1664V256zm0 512V640H0v128h1664z"/>`),
		g.Group(children),
	)
}
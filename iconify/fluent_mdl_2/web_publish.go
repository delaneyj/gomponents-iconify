package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebPublish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1024h-128V640H128v1024h1408v128H0V128h2048zm-128 128H128v256h1792V256zm-192 870l320 319l-91 91l-165-165v677h-128v-677l-165 165l-91-91l320-319z"/>`),
		g.Group(children),
	)
}
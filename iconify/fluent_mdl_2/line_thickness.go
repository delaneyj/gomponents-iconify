package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineThickness(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v128H0V384h2048zM0 768h2048v256H0V768zm0 512h2048v384H0v-384z"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SecondaryNav(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 384v128H0V384h1024zM0 896h2048v128H0V896zm1024 640v-128h1024v128H1024z"/>`),
		g.Group(children),
	)
}
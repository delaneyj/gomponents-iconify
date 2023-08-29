package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CollapseMenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 768V640h1792v128H128zm0-640h1792v128H128V128zm0 1152v-128h1792v128H128zm0 512v-128h1792v128H128z"/>`),
		g.Group(children),
	)
}
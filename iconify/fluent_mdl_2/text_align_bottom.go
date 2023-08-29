package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextAlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 896h1792v128H128V896zm256 512h1280v128H384v-128zm-256 640v-128h1792v128H128z"/>`),
		g.Group(children),
	)
}
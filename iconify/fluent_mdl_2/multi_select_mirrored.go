package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiSelectMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1294 384v128H14V384h1280zm0 512H14V768h1280v128zm0 384H14v-128h1280v128zm0 384H14v-128h1280v128zm400-388l286-286l68 68l-354 354l-178-178l68-68l110 110zm0-768l286-286l68 68l-354 354l-178-178l68-68l110 110z"/>`),
		g.Group(children),
	)
}
package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiSelect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v128H768V384h1280zM768 768h1280v128H768V768zm0 384h1280v128H768v-128zm0 384h1280v128H768v-128zM478 990l68 68l-354 354l-178-178l68-68l110 110l286-286zm0-768l68 68l-354 354L14 466l68-68l110 110l286-286z"/>`),
		g.Group(children),
	)
}
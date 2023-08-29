package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 384v-43h384v43H0zm128-85v-43h256v43H128zM0 213v-42h384v42H0zm128-85V85h256v43H128zM0 0h384v43H0V0z"/>`),
		g.Group(children),
	)
}
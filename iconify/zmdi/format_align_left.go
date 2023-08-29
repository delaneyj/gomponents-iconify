package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M256 256v43H0v-43h256zm0-171v43H0V85h256zM0 213v-42h384v42H0zm0 171v-43h384v43H0zM0 0h384v43H0V0z"/>`),
		g.Group(children),
	)
}
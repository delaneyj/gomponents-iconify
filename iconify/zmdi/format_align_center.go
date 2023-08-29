package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M85 256h214v43H85v-43zM0 384v-43h384v43H0zm0-171v-42h384v42H0zM85 85h214v43H85V85zM0 0h384v43H0V0z"/>`),
		g.Group(children),
	)
}
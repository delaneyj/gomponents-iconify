package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatIndentDecrease(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 299v-43h213v43H171zM0 192l85-85v170zm0 192v-43h384v43H0zM0 0h384v43H0V0zm171 128V85h213v43H171zm0 85v-42h213v42H171z"/>`),
		g.Group(children),
	)
}
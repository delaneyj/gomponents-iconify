package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M171 213L0 43V0h384v43L213 213v128h107v43H64v-43h107V213zM96 85h192l43-42H53z"/>`),
		g.Group(children),
	)
}